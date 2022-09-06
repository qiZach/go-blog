package server

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

var App = &ZsqServer{}

type ZsqServer struct {
}

func init() {
	// 模板加载到 common.Template
	common.LoadTemplate()
}
func (*ZsqServer) StartApplication(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
