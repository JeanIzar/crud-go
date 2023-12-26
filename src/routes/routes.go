package routes

import (
	"github.com/JeanIzar/crud-go/src/controllers"
	"github.com/gin-gonic/gin"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo", controllers.GetAllTodos)
	route.GET("/todo/:idTodo", controllers.GetIdTodo)
	route.PUT("/todo/:idTodo", controllers.UpdateTodo)
	route.DELETE("/todo/:idTodo", controllers.DeleteTodo)

	// Run route whenever triggered
	route.Run(":3000")
}
