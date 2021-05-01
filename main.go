package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/kushsharma/go-plug/transformer"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

func main() {
	fmt.Println("starting go-plug core application")

	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "go-plug core",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// handshakeConfigs are used to just do a basic handshake between
	// a plugin and host. If the handshake fails, a user friendly error is shown.
	// This prevents users from executing bad plugins or executing a plugin
	// directory. It is a UX feature, not a security feature.
	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "TRANSFORMER_PLUGIN",
		MagicCookieValue: "go-plug",
	}

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"transformer": &transformer.Plugin{},
	}

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugins/sql/main"),
		Logger:          logger,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via GRPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("transformer")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	transformerClient := raw.(transformer.Interface)

	if n, err := transformerClient.Name(); err != nil{
		panic(err)
	} else {
		fmt.Println("name of the plugin: ", n)
	}
	if d, err := transformerClient.Description(); err != nil{
		panic(err)
	} else {
		fmt.Println("description of the plugin: ", d)
	}
}
