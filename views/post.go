package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"net/http"
	"strconv"
	"strings"
)

// PostDetail 返回文章详情页
func (*HTMLApi) PostDetail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	// http://localhost:8080/p/1    1是参数，文章id
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("路径识别失败"))
		return
	}
	postRes, err := service.GetPostDetail(pId)
	if err != nil {
		detail.WriteError(w, errors.New("查询文章出错"))
		return
	}
	detail.WriteData(w, postRes)
}
