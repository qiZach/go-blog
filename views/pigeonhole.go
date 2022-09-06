package views

import (
	"go-blog/common"
	"go-blog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	// 加载页面
	pigeonhole := common.Template.Pigeonhole
	// 进行归档
	pigeonholeRes := service.FindPostPigeonhole()
	// 写回数据
	pigeonhole.WriteData(w, pigeonholeRes)
}
