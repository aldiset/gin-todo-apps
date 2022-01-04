package routes

import (
	"app/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/activity-groups", controllers.FindActivityGroups)
	r.POST("/activity-groups", controllers.CreateActivityGroups)
	r.GET("/activity-groups/:id", controllers.FindActivityGroup)
	r.PATCH("/activity-groups/:id", controllers.UpdateActivityGroups)
	r.DELETE("activity-groups/:id", controllers.DeleteActivityGroups)

	r.GET("/todo-items", controllers.FindToDos)
	r.POST("/todo-items", controllers.CreateToDo)
	r.GET("/todo-items/:id", controllers.FindToDo)
	r.PATCH("/todo-items/:id", controllers.UpdateToDo)
	r.DELETE("todo-items/:id", controllers.DeleteToDo)

	return r
}