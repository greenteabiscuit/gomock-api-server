package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/gomock-api-server/domain/model"
	"github.com/greenteabiscuit/gomock-api-server/infrastructure/database/entity"
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
	e := entity.Todo{Item: "Homework", IsDone: false}
	if err := db.Create(&e).Error; err != nil {
		return 0, err
	}
	return e.ID, nil
}

func (r *TodoRepository) Delete(
	ctx *gin.Context,
	db *gorm.DB,
) error {
	e := entity.Todo{ID: 1, Item: "Homework", IsDone: false}
	if err := db.Delete(&e).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) List(
	ctx *gin.Context,
	db *gorm.DB,
) ([]*model.Todo, error) {
	var todos []*entity.Todo
	var todoModels []*model.Todo
	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}
	for _, t := range todos {
		todoModels = append(todoModels, &model.Todo{ID: t.ID, Item: t.Item, IsDone: t.IsDone})
	}
	return todoModels, nil
}
