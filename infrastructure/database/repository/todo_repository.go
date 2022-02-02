package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TodoRepository ...
type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) Add(
	ctx *gin.Context,
	db *gorm.DB,
) (int, error) {
	return 0, nil
}
