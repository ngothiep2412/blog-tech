package articlemodel

import "errors"

var (
	ErrCannotCreateArticle = errors.New("cannot create article")
	ErrArticleNotFound     = errors.New("article not found")
	ErrCannotUpdateArticle = errors.New("cannot update article")
	ErrCannotDeleteArticle = errors.New("cannot delete article")
)
