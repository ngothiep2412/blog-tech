package articletagmodel

type ArticleTag struct {
	ArticleID int `json:"article_id" gorm:"column:article_id"`
	TagID     int `json:"tag_id" gorm:"column:tag_id"`
}

func (ArticleTag) TableName() string {
	return "article_tags"
}

type ArticleTagCreate struct {
	ArticleID int   `json:"article_id"`
	TagID     []int `json:"tag_id"`
}
