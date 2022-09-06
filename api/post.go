package api

import (
	"errors"
	"go-blog/common"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/service"
	"go-blog/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("路径识别失败"))
		return
	}
	post, err := dao.GetPostById(pId)
	if err != nil {
		common.Error(w, err)
	}
	common.Success(w, post)
}

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	// 获取用户id, 判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登陆已过期"))
		return
	}
	uid := claim.Uid
	method := r.Method
	switch method {
	// POST save
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := int(params["type"].(float64))
		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cId,
			UserId:     uid,
			ViewCount:  0,
			Type:       postType,
			Slug:       slug,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	// PUT update
	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		cId := int(params["type"].(float64))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := int(params["type"].(float64))
		pid := int(params["pid"].(float64))
		post := &models.Post{
			Pid:        pid,
			Title:      title,
			Content:    content,
			Markdown:   markdown,
			CategoryId: cId,
			UserId:     uid,
			ViewCount:  0,
			Type:       postType,
			Slug:       slug,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}

}
