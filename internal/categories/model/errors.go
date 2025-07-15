package categorymodel

import "errors"

var (
	ErrCategoryNotFound     = errors.New("category not found")
	ErrCategoryExists       = errors.New("category already exists")
	ErrCannotCreateCategory = errors.New("cannot create category")
	ErrCannotUpdateCategory = errors.New("cannot update category")
	ErrCannotDeleteCategory = errors.New("cannot delete category")
)
