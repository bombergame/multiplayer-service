package config

import (
	"github.com/bombergame/common/args"
	"github.com/bombergame/common/env"
)

var (
	ServiceName = "multiplayer-service"

	HttpPort = args.GetString("http_port", "80")

	RegistryAddress = env.GetVar("REGISTRY_ADDRESS", "127.0.0.1:8500")
)
