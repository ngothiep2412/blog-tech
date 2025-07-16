package categorymodel

import "blog-tech/common"

type Category struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;unique"`
	Slug            string `json:"slug" gorm:"column:slug;unique"`
	Description     string `json:"description" gorm:"column:description"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryCreate struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CategoryUpdate struct {
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
}
