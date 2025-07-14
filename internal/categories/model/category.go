package categorymodel

import "blog-tech/common"

type Category struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;unique"`
	Slug            string `json:"slug" gorm:"column:slug;unique"`
	Description     string `json:"description" gorm:"column:description"`
	Color           string `json:"color" gorm:"column:color"`
}
