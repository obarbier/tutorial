package main

import (
	"example-plugin/grpc/shared"
	"fmt"
	"github.com/hashicorp/go-plugin"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

//go:generate go build -o kv
func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("KV_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	pluginBuildPath := strings.Split(os.Getenv("KV_PLUGIN"), "/")
	pluginBuild := pluginBuildPath[len(pluginBuildPath)-1]

	dispenser := ""
	switch pluginBuild {
	case "kv-go-grpc":
		dispenser = "kv_grpc"
	case "kv-go-rpc":
		dispenser = "kv"
	default:
		log.Printf("unknown plugin :%v, using default", pluginBuild)
		dispenser = "kv_grpc"
	}
	// Request the plugin
	raw, err := rpcClient.Dispense(dispenser)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// We should have a KV store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	kv := raw.(shared.KV)
	os.Args = os.Args[1:]
	switch os.Args[0] {
	case "get":
		result, err := kv.Get(os.Args[1])
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(result))

	case "put":
		err := kv.Put(os.Args[1], []byte(os.Args[2]))
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}

	default:
		fmt.Printf("Please only use 'get' or 'put', given: %q", os.Args[0])
		os.Exit(1)
	}
	os.Exit(0)
}
