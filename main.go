package main

import (
	"github.com/i-coder-robot/go-micro-action-user/config"
	"github.com/i-coder-robot/go-micro-action-user/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	var Conf = config.MicroConfig{}
	service := micro.NewService(
		micro.Name(Conf.Name),
		micro.Version(Conf.Version),
	)
	service.Init()
	h := handler.Handler{
		Server: service.Server(),
	}
	h.Register()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
	log.Log("serviser run ...")

}
