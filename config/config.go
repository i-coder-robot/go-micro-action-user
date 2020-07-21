package config

import "github.com/micro/go-micro/config"

type MicroConfig struct {
	Name    string
	Version string
}

var Conf config.Config = MicroConfig{
	Name:    "go.micro.srv.user",
	Version: "latest",
}
