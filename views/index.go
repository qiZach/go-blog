package views

import (
	"go-blog/common"
	"go-blog/config"
	"go-blog/models"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
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
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		Viewer:     config.Cfg.Viewer,
		Categories: categories,
		Posts:      posts,
		Total:      1,
		Page:       1,
		Pages:      []int{1},
		PageEnd:    true,
	}
	index.WriteData(w, hr)
}
