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
		ctx *gin.Context, db *gorm.DB, ID int,
	) error
	AddTodos(
		ctx *gin.Context, db *gorm.DB, item string,
	) (int, error)
	UpdateTodos(
		ctx *gin.Context, db *gorm.DB, item string,
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

func (t *TodoUsecase) DeleteTodo(ctx *gin.Context, db *gorm.DB, ID int) error {
	return t.todoRepository.Delete(ctx, db, ID)
}

func (t *TodoUsecase) AddTodos(
	ctx *gin.Context, db *gorm.DB, item string,
) (int, error) {
	return t.todoRepository.Add(ctx, db, item)
}

func (t *TodoUsecase) UpdateTodos(
	ctx *gin.Context, db *gorm.DB, item string,
) error {
	return t.todoRepository.Update(ctx, db, item)
}