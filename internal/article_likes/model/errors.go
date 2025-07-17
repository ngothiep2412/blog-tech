package articlelikemodel

import "errors"

var (
	ErrCannotCreateArticleLike = errors.New("cannot create article like")
	ErrCannotDeleteArticleLike = errors.New("cannot delete article like")
	ErrCannotUpdateArticleLike = errors.New("cannot update article like")
	ErrArticleLikeNotFound     = errors.New("article like not found")
	ErrArticleLikeExisted      = errors.New("article like existed")
)
