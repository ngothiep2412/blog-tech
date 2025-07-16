package categorybiz

import (
	"blog-tech/common"
	categorymodel "blog-tech/internal/categories/model"
	categoryrepository "blog-tech/internal/categories/repository/mysql"
	usermodel "blog-tech/internal/users/model"
	"context"
)

type CategoryBusiness interface {
	CreateCategory(ctx context.Context, userID int, category *categorymodel.CategoryCreate) error
	UpdateCategory(ctx context.Context, category *categorymodel.CategoryUpdate) error
	GetCategoryByID(ctx context.Context, id int) (*categorymodel.Category, error)
}

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (usermodel.User, error)
}

type categoryBusiness struct {
	categoryRepo categoryrepository.CategoryRepository
	userRepo     UserRepository
}

func NewCategoryBusiness(categoryRepo categoryrepository.CategoryRepository, userRepo UserRepository) *categoryBusiness {
	return &categoryBusiness{
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (b *categoryBusiness) CreateCategory(ctx context.Context, userID int, category *categorymodel.CategoryCreate) error {

	user, err := b.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}

	if user.ID == 0 {
		return usermodel.ErrUserNotFound
	}

	err = b.categoryRepo.CreateCategory(ctx, category)

	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}
	return nil
}

func (b *categoryBusiness) UpdateCategory(ctx context.Context, category *categorymodel.CategoryUpdate) error {
	err := b.categoryRepo.UpdateCategory(ctx, category)
	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}
	return nil
}

func (b *categoryBusiness) GetCategoryByID(ctx context.Context, id int) (*categorymodel.Category, error) {
	category, err := b.categoryRepo.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return category, nil
}
