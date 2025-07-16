package tagbiz

import (
	"blog-tech/common"
	tagmodel "blog-tech/internal/tags/model"
	usermodel "blog-tech/internal/users/model"
	"context"
)

type TagRepository interface {
	CreateTag(ctx context.Context, input *tagmodel.TagCreate) (*tagmodel.Tag, error)
	GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error)
	UpdateTag(ctx context.Context, id int, tag *tagmodel.TagUpdate) error
	GetTagByName(ctx context.Context, name string) (*tagmodel.Tag, error)
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

func (b *business) CreateTag(ctx context.Context, userID int, tag *tagmodel.TagCreate) (*tagmodel.Tag, error) {
	user, err := b.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}

	if user.ID == 0 {
		return nil, usermodel.ErrUserNotFound
	}

	tagResp, err := b.tagRepo.CreateTag(ctx, tag)

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return tagResp, nil
}

func (b *business) CreateTagRPC(ctx context.Context, tag *tagmodel.TagCreate) (*tagmodel.Tag, error) {
	tagResp, err := b.tagRepo.CreateTag(ctx, tag)

	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return tagResp, nil
}

func (b *business) GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error) {
	tag, err := b.tagRepo.GetTagByID(ctx, id)
	if err != nil {
		return nil, common.ErrInternalServerError.WithError(err.Error())
	}
	return tag, nil
}

func (b *business) GetTagByName(ctx context.Context, name string) (*tagmodel.Tag, error) {
	tag, err := b.tagRepo.GetTagByName(ctx, name)
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
