package tagrepomysql

import (
	tagmodel "blog-tech/internal/tags/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *tagRepository {
	return &tagRepository{db: db}
}

func (s *tagRepository) CreateTag(ctx context.Context, input *tagmodel.TagCreate) (*tagmodel.Tag, error) {
	tag := &tagmodel.Tag{
		Name: input.Name,
		Slug: input.Slug,
	}

	err := s.db.Create(tag).Error
	if err != nil {
		return nil, errors.Wrap(err, tagmodel.ErrCannotCreateTag.Error())
	}

	return tag, nil
}

func (s *tagRepository) GetTagByID(ctx context.Context, id int) (*tagmodel.Tag, error) {
	var tag tagmodel.Tag

	if err := s.db.Table(tagmodel.Tag{}.TableName()).Where("id = ?", id).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, tagmodel.ErrTagNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &tag, nil
}

func (s *tagRepository) GetTagByName(ctx context.Context, name string) (*tagmodel.Tag, error) {
	var tag tagmodel.Tag

	if err := s.db.Table(tagmodel.Tag{}.TableName()).Where("name = ?", name).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, tagmodel.ErrTagNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}

	return &tag, nil
}

func (s *tagRepository) UpdateTag(ctx context.Context, id int, tag *tagmodel.TagUpdate) error {
	if err := s.db.Table(tagmodel.Tag{}.TableName()).Where("id = ?", id).Updates(tag).Error; err != nil {
		return errors.Wrap(err, tagmodel.ErrCannotUpdateTag.Error())
	}

	return nil
}
