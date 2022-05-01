package main

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"golang.org/x/exp/maps"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type keyType int

const (
	ED25519 keyType = iota
	RSA
)

var KEYTYPE = map[string]keyType{
	"RSA":     RSA,
	"ED25519": ED25519,
}
var reader = rand.Reader
var wg sync.WaitGroup
var mr sync.Mutex
var number = flag.Int("n", 1, "How many keys you want to create")
var baseName = flag.String("b", "", "What is the base name of the key")
var kType = flag.String("kt", "ED25519", fmt.Sprintf("What type of keys: choice are (%s)", strings.Join(maps.Keys(KEYTYPE), ",")))

func main() {
	//interruptChan := make(chan os.Signal, 1)
	flag.Parse()
	if *baseName == "" {
		log.Fatal("-b need to be set")
	}
	if _, ok := KEYTYPE[*kType]; !ok {
		log.Fatalf("%s is not a valid keyType", *kType)
	}
	for i := 1; i <= *number; i++ {
		wg.Add(1)
		go func(idx int, baseName string) {
			defer wg.Done()
			f1, err := os.Create(fmt.Sprintf("%d-%s.priv", idx, baseName))
			if err != nil {
				log.Fatal(fmt.Errorf("failed to create File: %v", err))
			}
			f2, err := os.Create(fmt.Sprintf("%d-%s.pub", idx, baseName))
			if err != nil {
				log.Fatal(fmt.Errorf("failed to create File: %v", err))
			}
			generate(KEYTYPE[*kType], f1, f2)
			f1.Close()
			f2.Close()

		}(i, *baseName)
	}
	wg.Wait()

	//signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
}
func generate(kType keyType, f1, f2 io.Writer) {
	switch kType {
	case ED25519:
		generateEd25519(f1, f2)
	case RSA:
		generateRsa(f1, f2)
	default:
		log.Fatal("keyType not yet implemented")
	}
}

func generateRsa(f1, f2 io.Writer) {
	mr.Lock()
	pk, err := rsa.GenerateKey(reader, 2048)
	mr.Unlock()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to generate key-pair: %v", err))
	}
	err = pem.Encode(f1, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk),
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to encode key-pair: %v", err))
	}
	err = pem.Encode(f2, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&pk.PublicKey),
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to encode key-pair: %v", err))
	}

	if err != nil {
		log.Fatal(fmt.Errorf("failed to close file: %v", err))
	}

}

func generateEd25519(f1, f2 io.Writer) {
	mr.Lock()
	pub, pk, err := ed25519.GenerateKey(reader)
	mr.Unlock()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to generate key-pair: %v", err))
	}
	err = pem.Encode(f1, &pem.Block{
		Type:  "ED25519 PRIVATE KEY",
		Bytes: pk,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to encode key-pair: %v", err))
	}
	err = pem.Encode(f2, &pem.Block{
		Type:  "ED25519 PUBLIC KEY",
		Bytes: pub,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("failed to encode key-pair: %v", err))
	}

	if err != nil {
		log.Fatal(fmt.Errorf("failed to close file: %v", err))
	}

}

func Sign(keyType keyType, signer []byte, message []byte) string {
	switch keyType {
	case ED25519:
		return signEd(signer, message)
	case RSA:
		return signRsa(signer, message)
	default:
		log.Fatal("keyType is not yet implemented")
		return ""
	}
}

func signEd(privateKey ed25519.PrivateKey, message []byte) string {
	return string(ed25519.Sign(privateKey, message))
}

func signRsa(privateKey []byte, message []byte) string {
	priv, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		log.Fatalf("failed to parse pk: %s", err)
	}

	msgHash := sha512.New()
	_, err = msgHash.Write(message)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	mr.Lock()
	b, err := rsa.SignPKCS1v15(reader, priv, crypto.SHA512, msgHashSum)
	mr.Unlock()
	if err != nil {
		log.Fatalf("failed to sign hash: %s", err)
	}
	return string(b)
}

func Verify(keyType keyType, publicKey []byte, message, sig []byte) bool {
	switch keyType {
	case ED25519:
		return verifyEd(publicKey, message, sig)
	case RSA:
		return verifyRsa(publicKey, message, sig)
	default:
		log.Fatal("keyType is not yet implemented")
		return false
	}
}

func verifyEd(publicKey ed25519.PublicKey, message, sig []byte) bool {
	return ed25519.Verify(publicKey, message, sig)
}

func verifyRsa(publicKey []byte, message, sig []byte) bool {
	privateKeyBytes, err := x509.ParsePKCS1PublicKey(publicKey)
	if err != nil {
		log.Fatalf("failed to verify hash: %s", err)
	}
	msgHash := sha512.New()
	_, err = msgHash.Write(message)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)
	err = rsa.VerifyPKCS1v15(privateKeyBytes, crypto.SHA512, msgHashSum, sig)
	if err != nil {
		log.Fatalf("failed to verify: %s", err)
		return false
	}
	return true
}
