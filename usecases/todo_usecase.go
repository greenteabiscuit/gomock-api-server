package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/gomock-api-server/domain/repository"
)

type TodoUsecaseInterface interface {
	FindAllTodos(
		ctx *gin.Context,
	) ()
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

func (t *TodoUsecase) FindAllTodos(ctx *gin.Context){
	return
}

func (t *TodoUsecase) DeleteTodo(ctx *gin.Context) error {
	return nil
}

func (t *TodoUsecase) AddTodos(
	ctx *gin.Context,
) error {
	return nil
}
