package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GormComponent interface {
	GetDB() *gorm.DB
}

type ConfigComponent interface {
	GetGRPCPort() int
	GetGRPCServerAddress() string
	GetGRPCUserAddress() string
	GetGRPCCategoryAddress() string
}

type GINComponent interface {
	GetPort() int
	GetRouter() *gin.Engine
}
