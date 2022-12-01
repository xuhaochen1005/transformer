package todolist

import (
	"github.com/gin-gonic/gin"
	"transformer_mz/api/todolist"
)

func SetupTodoListGroup(rg *gin.RouterGroup) {
	todoListRouter := rg.Group("/todolist")

	// get todolist list
	todoListRouter.GET("/", todolist.FindTodoLists)
	// create todoList
	todoListRouter.POST("/", todolist.CreateTodoList)
	// update todolist
	todoListRouter.PATCH("/:id", todolist.UpdateTodoList)
	// delete todolist
	todoListRouter.DELETE("/:id", todolist.DeleteTodoList)
}
