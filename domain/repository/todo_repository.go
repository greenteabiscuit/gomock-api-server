package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Add(
		ctx *gin.Context,
		db *gorm.DB,
	) (int, error)
}
