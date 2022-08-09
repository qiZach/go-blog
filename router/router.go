package router

import (
	"go-blog/api"
	"go-blog/views"

	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func Router() {
	// 1.页面 2.数据(json) 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)

	// 启动静态文件服务
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))
}
