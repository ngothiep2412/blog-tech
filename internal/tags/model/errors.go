package tagmodel

import "errors"

var (
	ErrCannotCreateTag = errors.New("cannot create tag")
	ErrTagNotFound     = errors.New("tag not found")
	ErrCannotUpdateTag = errors.New("cannot update tag")
)
