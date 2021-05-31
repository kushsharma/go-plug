package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/kushsharma/go-plug/transformer"
)

type SQLTransform struct {
}

func (B SQLTransform) Name() (string, error) {
	return "SQL", nil
}

func (B SQLTransform) Description() (string, error) {
	return "test sql transformer", nil
}

func (B SQLTransform) GenerateDependencies(s string) ([]string, error) {
	return []string{"test"}, nil
}

func main() {
	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "TRANSFORMER_PLUGIN",
		MagicCookieValue: "go-plug",
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"transformer": &transformer.Plugin{Impl: &SQLTransform{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
