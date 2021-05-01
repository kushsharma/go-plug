package transformer

import (
	"context"
	"github.com/kushsharma/go-plug/proto"
)

// GRPCClient will be used by core to talk over grpc with plugins
type GRPCClient struct {
	client proto.TransformerClient
}

func (m *GRPCClient) Name() (string, error) {
	resp, err := m.client.GetName(context.Background(), &proto.GetNameRequest{})
	if err != nil {
		return "", err
	}
	return resp.Name, nil
}

func (m *GRPCClient) Description() (string, error) {
	resp, err := m.client.GetDescription(context.Background(), &proto.GetDescriptionRequest{})
	if err != nil {
		return "", err
	}
	return resp.Description, nil
}

func (m *GRPCClient) GenerateDependencies(request string) ([]string, error) {
	resp, err := m.client.GenerateDependencies(context.Background(), &proto.GenerateDependenciesRequest{
		Random: request,
	})
	if err != nil {
		return nil, err
	}
	return resp.Dependencies, nil
}

// GRPCServer will be used by plugins
// this is working as proto adapter
type GRPCServer struct{
	// This is the real implementation
	Impl Interface

	proto.UnimplementedTransformerServer
}

func (m *GRPCServer) GetName(ctx context.Context, req *proto.GetNameRequest) (*proto.GetNameResponse, error) {
	n, err := m.Impl.Name()
	if err != nil {
		return nil, err
	}
	return &proto.GetNameResponse{Name: n}, nil
}

func (m *GRPCServer) GetDescription(ctx context.Context, req *proto.GetDescriptionRequest) (*proto.GetDescriptionResponse, error) {
	d, err := m.Impl.Description()
	if err != nil {
		return nil, err
	}
	return &proto.GetDescriptionResponse{Description: d}, nil
}

func (m *GRPCServer) GenerateDependencies(ctx context.Context, req *proto.GenerateDependenciesRequest) (*proto.GenerateDependenciesResponse, error) {
	d, err := m.Impl.GenerateDependencies(req.Random)
	if err != nil {
		return nil, err
	}
	return &proto.GenerateDependenciesResponse{Dependencies: d}, nil
}
