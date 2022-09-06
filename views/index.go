package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	// 拿到Index页面
	index := common.Template.Index
	// 页面上涉及到的所有的数据，必须有定义，给页面填入数据
	// 数据库查询
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败: ", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	// 每页显示的数量
	pageSize := 10
	hr, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Index获取数据出错: ", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员"))
	}
	// 返回，发送至浏览器解析
	index.WriteData(w, hr)
}
