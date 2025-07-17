package articlelikemodel

import "blog-tech/common"

const EntityName = "ArticleLike"

type ArticleLike struct {
	common.SqlModel `json:",inline"`
	ArticleID       int `json:"article_id" gorm:"column:article_id"`
	UserID          int `json:"user_id" gorm:"column:user_id"`
}

func (ArticleLike) TableName() string {
	return "article_likes"
}
