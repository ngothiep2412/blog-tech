package articletagmodel

import "blog-tech/common"

type ArticleTag struct {
	common.SqlModel `json:",inline"`
	ArticleID       uint `json:"article_id" gorm:"column:article_id"`
	TagID           uint `json:"tag_id" gorm:"column:tag_id"`
}
