package transformer

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/kushsharma/go-plug/proto"
	"google.golang.org/grpc"
)

type Plugin struct {
	plugin.NetRPCUnsupportedPlugin
	plugin.GRPCPlugin

	Impl Interface
}

func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterTransformerServer(s, &GRPCServer{
		Impl:   p.Impl,
	})
	return nil
}

func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewTransformerClient(c),
	}, nil
}

var _ plugin.GRPCPlugin = &Plugin{}