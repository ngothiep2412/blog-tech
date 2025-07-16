package articlemodel

import (
	"blog-tech/common"
	tagmodel "blog-tech/internal/tags/model"
	"time"
)

type Article struct {
	common.SqlModel  `json:",inline"`
	UserID           int        `json:"user_id" gorm:"column:user_id"`
	CategoryID       int        `json:"category_id" gorm:"column:category_id"`
	Title            string     `json:"title" gorm:"column:title"`
	Slug             string     `json:"slug" gorm:"column:slug"`
	Content          string     `json:"content" gorm:"column:content"`
	Excerpt          string     `json:"excerpt" gorm:"column:excerpt"`
	FeaturedImageURL string     `json:"featured_image_url" gorm:"column:featured_image_url"`
	Status           string     `json:"status" gorm:"column:status"`
	ViewCount        int        `json:"view_count" gorm:"column:view_count"`
	LikeCount        int        `json:"like_count" gorm:"column:like_count"`
	ShareCount       int        `json:"share_count" gorm:"column:share_count"`
	PublishedAt      *time.Time `json:"published_at" gorm:"column:published_at"`

	Tags   []tagmodel.Tag `json:"tags,omitempty" gorm:"-"`
	TagIDs []int          `json:"tag_ids,omitempty" gorm:"-"`
}

func (Article) TableName() string {
	return "articles"
}

type ArticleCreate struct {
	common.SqlModel  `json:",inline"`
	UserID           int    `json:"user_id" validate:"required,gt=0"`
	CategoryID       int    `json:"category_id" validate:"required,gt=0"`
	Title            string `json:"title" validate:"required,min=1,max=255"`
	Content          string `json:"content" validate:"required,min=10"`
	Excerpt          string `json:"excerpt" validate:"max=500"`
	FeaturedImageURL string `json:"featured_image_url" validate:"omitempty,url"`
	Status           string `json:"status" validate:"required,oneof=draft published archived"`
	Slug             string `json:"slug" gorm:"column:slug"`

	Tags []tagmodel.Tag `json:"tags" gorm:"-"`
}
