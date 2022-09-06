package main

import (
	"go-blog/config"
	"go-blog/server"
)

func main() {
	server.App.StartApplication(config.Cfg.Server.Ip, config.Cfg.Server.Port)
}
