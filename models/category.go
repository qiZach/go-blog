package models

type Category struct {
	Cid      int    `orm:"cid" json:"cid"`
	Name     string `orm:"name" json:"name"`
	CreateAt string `orm:"create_at" json:"createAt"`
	UpdateAt string `orm:"update_at" json:"updateAt"`
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
