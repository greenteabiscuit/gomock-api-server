package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/gomock-api-server/usecases"
	"gorm.io/gorm"
	"net/http"
)

type TodoControllerInterface interface {
	List()
	Delete()
	Add()
}

type TodoController struct {
	db          *gorm.DB
	todoUsecase usecases.TodoUsecaseInterface
}

type TodoResponse struct {
	ID     int    `json:"id"`
	Item   string `json:"item"`
	IsDone bool   `json:"is_done"`
}

type TodoListResponse struct {
	TodoList []TodoResponse
}

func NewTodoController(db *gorm.DB, todoUsecaseInterface usecases.TodoUsecaseInterface) *TodoController {
	return &TodoController{db: db, todoUsecase: todoUsecaseInterface}
}

func (t *TodoController) List(ctx *gin.Context) {
	models, err := t.todoUsecase.FindAllTodos(ctx, t.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "bad request")
		return
	}
	var lis []TodoResponse
	var res TodoListResponse
	for _, m := range models {
		lis = append(lis, TodoResponse{ID: m.ID, Item: m.Item, IsDone: m.IsDone})
	}
	res.TodoList = lis
	ctx.JSON(http.StatusOK, res)
}

type DeleteRequest struct {
	ID int `json: "id"`
}

func (t *TodoController) Delete(ctx *gin.Context) {
	var dr DeleteRequest
	ctx.BindJSON(&dr)
	if err := t.todoUsecase.DeleteTodo(ctx, t.db, dr.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, "bad request")
		return
	}
	ctx.JSON(http.StatusOK, "delete success")
}

type AddRequest struct {
	Item string `json: "item"`
}

func (t *TodoController) Add(ctx *gin.Context) {
	var ar AddRequest
	ctx.BindJSON(&ar)
	ID, err := t.todoUsecase.AddTodos(ctx, t.db, ar.Item)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "bad request")
		return
	}
	ctx.JSON(http.StatusOK, ID)
}
