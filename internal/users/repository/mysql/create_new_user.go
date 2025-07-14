package usermysql

import (
	usermodel "blog-tech/internal/users/model"
	"context"

	"github.com/pkg/errors"
)

func (repo *mysqlStore) CreateNewUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.Wrap(err, usermodel.ErrCannotCreateUser.Error())
	}
	return nil
}
