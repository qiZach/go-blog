package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	// 查询所有的文章,进行月份的整理
	posts, _ := dao.GetAllPost()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	// 查询所有的分类
	category, _ := dao.GetAllCategory()
	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categories:   category,
		Lines:        pigeonholeMap,
	}
}
