package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

func init() {
	// 模板加载到 common.Template
	common.LoadTemplate()
}

func main() {

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
