package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "zsqgo")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	// 组装loginResult
	// 1. 生成Token jwt技术进行生成令牌
	uid := user.Uid
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo = models.UserInfo{
		Uid:      user.Uid,
		UserName: user.UserName,
		Avatar:   user.Avatar,
	}
	var lr = &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}
	return lr, nil
}
