package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/gomock-api-server/usecases"
	"gorm.io/gorm"
)

type TodoControllerInterface interface {
	List()
	Delete()
	Add()
}

type TodoController struct {
	db                  *gorm.DB
	todoUsecase usecases.TodoUsecaseInterface
}

func NewTodoController(db *gorm.DB, todoUsecaseInterface usecases.TodoUsecaseInterface) *TodoController {
	return &TodoController{db: db, todoUsecase: todoUsecaseInterface}
}

func (t *TodoController) List(ctx *gin.Context) {}

func (t *TodoController) Delete(ctx *gin.Context) {}

func (t *TodoController) Add(ctx *gin.Context) {}
