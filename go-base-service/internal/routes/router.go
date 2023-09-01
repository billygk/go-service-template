package routes

import (
	"github.com/billygk/go-service-template/go-base-service/internal/controllers"
	"github.com/billygk/go-service-template/go-base-service/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRouter sets up the router accept db as a parameter
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	todoRepository := repositories.NewTodoRepository(db)
	todoController := controllers.NewTodoController(todoRepository)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/todos", todoController.Create)
		v1.GET("/todos", todoController.FindAll)
		v1.GET("/todos/:id", todoController.FindByID)
		v1.PUT("/todos/:id", todoController.Update)
		v1.DELETE("/todos/:id", todoController.Delete)
	}

	return router
}
