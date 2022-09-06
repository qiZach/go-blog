package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	// 返回三种类型： 1.页面  2.  api数据(json) 3.静态资源
	// 返回主页，并加载主页文章 views.HTML.Index 为处理函数
	http.HandleFunc("/", views.HTML.Index)
	// 按照分类返回文章
	http.HandleFunc("/c/", views.HTML.Category)
	// 展示文章详情
	http.HandleFunc("/p/", views.HTML.PostDetail)
	// 加载登录页面
	http.HandleFunc("/login", views.HTML.Login)
	// 加载写作页面
	http.HandleFunc("/writing", views.HTML.Writing)
	// 加载归档页面
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	// 处理登录请求
	http.HandleFunc("/api/v1/login", api.API.Login)
	// 处理Post 保存文章或更新文章
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	// 根据pid 获取文章
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	// 根据搜索查询文章
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	//  获取qiniu token
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)

	// 启动静态文件服务
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource"))))
}
