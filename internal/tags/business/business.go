package tagbiz

import (
	"blog-tech/common"
	tagmodel "blog-tech/internal/tags/model"
	usermodel "blog-tech/internal/users/model"
	"context"
)

type TagRepository interface {
	CreateTag(ctx context.Context, tag *tagmodel.TagCreate) error
	GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error)
	UpdateTag(ctx context.Context, id int, tag *tagmodel.TagUpdate) error
}

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (*usermodel.User, error)
}

type business struct {
	tagRepo  TagRepository
	userRepo UserRepository
}

func NewBusiness(tagRepo TagRepository, userRepo UserRepository) *business {
	return &business{
		tagRepo:  tagRepo,
		userRepo: userRepo,
	}
}

func (b *business) CreateTag(ctx context.Context, userID int, tag *tagmodel.TagCreate) error {
	user, err := b.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}

	if user.ID == 0 {
		return usermodel.ErrUserNotFound
	}

	err = b.tagRepo.CreateTag(ctx, tag)
	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}
	return nil
}

func (b *business) GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error) {
	tag, err := b.tagRepo.GetTagByID(ctx, id)
	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return tag, nil
}

func (b *business) UpdateTag(ctx context.Context, id int, tag *tagmodel.TagUpdate) error {
	err := b.tagRepo.UpdateTag(ctx, id, tag)
	if err != nil {
		return common.ErrInternalServerError.WithError(err.Error())
	}
	return nil
}
