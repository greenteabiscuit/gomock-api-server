package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/gomock-api-server/domain/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Add(
		ctx *gin.Context,
		db *gorm.DB,
		item string,
	) (int, error)
	Update(
		ctx *gin.Context,
		db *gorm.DB,
		item string,
	) error
	Delete(ctx *gin.Context, db *gorm.DB, ID int) error
	List(ctx *gin.Context, db *gorm.DB) ([]*model.Todo, error)
}
