package tagmodel

import "blog-tech/common"

type Tag struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;unique"`
	Slug            string `json:"slug" gorm:"column:slug;unique"`
}

func (Tag) TableName() string {
	return "tags"
}

type TagCreate struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type TagUpdate struct {
	Name *string `json:"name"`
	Slug *string `json:"slug"`
}
