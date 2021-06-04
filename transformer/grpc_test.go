package transformer

import (
	"context"
	"google.golang.org/grpc"
	"testing"

	"github.com/kushsharma/go-plug/proto"
	"github.com/stretchr/testify/mock"
)

type mockedClient struct {
	mock.Mock
}

func (c *mockedClient) GetName(ctx context.Context, in *proto.GetNameRequest, opts ...grpc.CallOption) (*proto.GetNameResponse, error) {
	args := c.Called(ctx, in)
	return args.Get(0).(*proto.GetNameResponse), args.Error(1)
}

func (c *mockedClient) GetDescription(ctx context.Context, in *proto.GetDescriptionRequest, opts ...grpc.CallOption) (*proto.GetDescriptionResponse, error) {
	args := c.Called(ctx, in)
	return args.Get(0).(*proto.GetDescriptionResponse), args.Error(1)
}

func (c *mockedClient) GenerateDependencies(ctx context.Context, in *proto.GenerateDependenciesRequest, opts ...grpc.CallOption) (*proto.GenerateDependenciesResponse, error) {
	args := c.Called(ctx, in)
	return args.Get(0).(*proto.GenerateDependenciesResponse), args.Error(1)
}

func TestGRPCClient_Name(t *testing.T) {
	client := &mockedClient{}
	client.On("GetName", context.Background(), &proto.GetNameRequest{}).Return(&proto.GetNameResponse{Name: "hello"}, nil)

	type fields struct {
		client proto.TransformerClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name:    "should return name correctly",
			fields:  fields{
				client: client,
			},
			want:    "hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GRPCClient{
				client: tt.fields.client,
			}
			got, err := m.Name()
			if (err != nil) != tt.wantErr {
				t.Errorf("GRPCClient.Name() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GRPCClient.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
