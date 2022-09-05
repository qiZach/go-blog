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
	// 返回三种类型： 1.页面  2.  api数据(json) 3.静态资源
	// views.HTML.Index 为处理函数
	http.HandleFunc("/", views.HTML.Index)
	// 处理Post请求
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)

	// 启动静态文件服务
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))
}
