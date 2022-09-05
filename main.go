package main

import (
	"embed"
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

func init() {
	// 模板加载到 common.Template
	common.LoadTemplate()
}

//go:embed public/resource/*
//go:embed template/*
var f embed.FS

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
