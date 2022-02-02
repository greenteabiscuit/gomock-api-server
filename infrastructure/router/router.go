package router

import (
	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/gomock-api-server/controllers"
	"github.com/greenteabiscuit/gomock-api-server/infrastructure/database/repository"
	"github.com/greenteabiscuit/gomock-api-server/usecases"
	"gorm.io/gorm"
)

// Router ...
type Router struct {
}

// InitRouter ...
func (router *Router) InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	todoController := controllers.NewTodoController(
		db,
		usecases.NewTodoUsecase(
			repository.NewTodoRepository(),
		),
	)

	version := "/v1"
	// 認証ルーター
	v1 := r.Group(version)
	{
		// アルバム
		todo := v1.Group("/todo")
		{
			todo.GET("/list", todoController.List)
			todo.DELETE("/delete", todoController.Delete)
			todo.POST("/add", todoController.Add)
		}

	}
	return r
}