package usermysql

import "gorm.io/gorm"

type mysqlStore struct {
	db *gorm.DB
}

func NewMysqlStore(db *gorm.DB) *mysqlStore {
	return &mysqlStore{db: db}
}
