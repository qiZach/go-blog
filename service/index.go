package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

func GetAllIndexInfo(slug string, page, pageSize int) (*models.HomeResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	var total int
	if slug == "" {
		posts, err = dao.GetPostPage(page, pageSize)
		total = dao.CountGetAllPost()
	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pageSize)
		total = dao.CountGetAllPostBySlug(slug)
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	// pages = (n-1)/10 + 1

	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	// 需要写入页面的数据
	var hr = &models.HomeResponse{
		Viewer:     config.Cfg.Viewer, // 视图内的基本数据
		Categories: categories,
		Posts:      postMores,
		Total:      total,
		Page:       page,  // 当前第几页
		Pages:      pages, // 显示第几页
		PageEnd:    page != pagesCount,
	}
	return hr, nil
}
