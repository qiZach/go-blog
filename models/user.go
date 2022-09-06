package models

import "time"

type User struct {
	Uid      int       `orm:"uid" json:"uid"`
	UserName string    `orm:"user_name" json:"userName"`
	Passwd   string    `orm:"passwd" json:"passwd"`
	Avatar   string    `orm:"avatar" json:"avatar"`
	CreateAt time.Time `orm:"create_at" json:"createAt"`
	UpdateAt time.Time `orm:"update_at" json:"updateAt"`
}

type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
