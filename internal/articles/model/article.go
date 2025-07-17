package articlemodel

import (
	"blog-tech/common"
	tagmodel "blog-tech/internal/tags/model"
	"time"
)

const (
	StatusDraft     = "draft"
	StatusPublished = "published"
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
	UserID           int        `json:"user_id" gorm:"column:user_id"`
	CategoryID       int        `json:"category_id" gorm:"column:category_id"`
	Title            string     `json:"title" gorm:"column:title"`
	Content          string     `json:"content" gorm:"column:content"`
	Excerpt          string     `json:"excerpt" gorm:"column:excerpt"`
	FeaturedImageURL string     `json:"featured_image_url" gorm:"column:featured_image_url"`
	Status           string     `json:"status" gorm:"column:status"`
	Slug             string     `json:"slug" gorm:"column:slug"`
	PublishedAt      *time.Time `json:"published_at" gorm:"column:published_at"`

	Tags []tagmodel.Tag `json:"tags" gorm:"-"`
}
