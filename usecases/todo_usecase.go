package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/gomock-api-server/domain/model"
	"github.com/greenteabiscuit/gomock-api-server/domain/repository"
	"gorm.io/gorm"
)

type TodoUsecaseInterface interface {
	FindAllTodos(
		ctx *gin.Context, db *gorm.DB,
	) ([]*model.Todo, error)
	DeleteTodo(
		ctx *gin.Context,
	) error
	AddTodos(
		ctx *gin.Context,
	) error
}

// TodoUsecase is
type TodoUsecase struct {
	todoRepository repository.TodoRepository
}

// NewTodoUsecase is
func NewTodoUsecase(todoRepository repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepository: todoRepository,
	}
}

func (t *TodoUsecase) FindAllTodos(ctx *gin.Context, db *gorm.DB) ([]*model.Todo, error) {
	return t.todoRepository.List(ctx, db)
}

func (t *TodoUsecase) DeleteTodo(ctx *gin.Context) error {
	return nil
}

func (t *TodoUsecase) AddTodos(
	ctx *gin.Context,
) error {
	return nil
}
