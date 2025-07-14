package articlelikemodel

import "blog-tech/common"

type ArticleLike struct {
	common.SqlModel `json:",inline"`
	ArticleID       uint `json:"article_id" gorm:"column:article_id"`
	UserID          uint `json:"user_id" gorm:"column:user_id"`
}
