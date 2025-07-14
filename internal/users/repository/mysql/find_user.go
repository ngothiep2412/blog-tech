package usermysql

import (
	usermodel "blog-tech/internal/users/model"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *mysqlStore) FindUser(ctx context.Context, condition map[string]interface{}) (*usermodel.User, error) {
	var user usermodel.User

	if err := repo.db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrap(err, usermodel.ErrCannotGetUser.Error())
		}
	}
	return &user, nil
}
