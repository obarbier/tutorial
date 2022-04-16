package shared

import (
	"context"
	proto "example-plugin/grpc/proto"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct{ client proto.KVClient }

func (m *GRPCClient) Put(key string, value []byte) error {
	_, err := m.client.Put(context.Background(), &proto.PutRequest{
		Key:   key,
		Value: value,
	})
	return err
}

func (m *GRPCClient) Get(key string) ([]byte, error) {
	resp, err := m.client.Get(context.Background(), &proto.GetRequest{
		Key: key,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	proto.UnimplementedKVServer
	// This is the real implementation
	Impl KV
}

func (m *GRPCServer) Get(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error) {
	v, err := m.Impl.Get(request.Key)
	return &proto.GetResponse{Value: v}, err
}

func (m *GRPCServer) Put(ctx context.Context, request *proto.PutRequest) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Put(request.Key, request.Value)
}
