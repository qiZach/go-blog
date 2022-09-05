package views

import (
	"go-blog/common"
	"go-blog/config"
	"go-blog/models"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	// 拿到Index页面
	index := common.Template.Index
	// 页面上涉及到的所有的数据，必须有定义，给页面填入数据
	// 数据库查询
	var categories = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "张思祺",
			ViewCount:    123,
			CreateAt:     "2022-09-02",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	// 需要写入页面的数据
	var hr = &models.HomeResponse{
		Viewer:     config.Cfg.Viewer,
		Categories: categories,
		Posts:      posts,
		Total:      1,
		Page:       1,
		Pages:      []int{1},
		PageEnd:    true,
	}
	// 返回，发送至浏览器解析
	index.WriteData(w, hr)
}
