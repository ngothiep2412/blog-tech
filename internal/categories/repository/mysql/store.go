package categorymysql

import (
	categorymodel "blog-tech/internal/categories/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *categorymodel.CategoryCreate) error
	UpdateCategory(ctx context.Context, category *categorymodel.CategoryUpdate) error
	GetCategoryByID(ctx context.Context, id int) (*categorymodel.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (s *categoryRepository) CreateCategory(ctx context.Context, category *categorymodel.CategoryCreate) error {
	if err := s.db.Table(categorymodel.Category{}.TableName()).Create(&category).Error; err != nil {
		return errors.Wrap(err, categorymodel.ErrCannotCreateCategory.Error())
	}
	return nil
}

func (s *categoryRepository) UpdateCategory(ctx context.Context, category *categorymodel.CategoryUpdate) error {
	if err := s.db.Table(categorymodel.Category{}.TableName()).Save(category).Error; err != nil {
		return errors.Wrap(err, categorymodel.ErrCannotUpdateCategory.Error())
	}
	return nil
}

func (s *categoryRepository) GetCategoryByID(ctx context.Context, id int) (*categorymodel.Category, error) {
	var category categorymodel.Category

	if err := s.db.Table(categorymodel.Category{}.TableName()).First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(err, categorymodel.ErrCategoryNotFound.Error())
		}
		return nil, errors.WithStack(err)
	}
	return &category, nil
}
