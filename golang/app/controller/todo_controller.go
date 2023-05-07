package controller

import (
	"example/todoapp/model"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	FetchTodos(c *gin.Context)
	AddTodo(c *gin.Context)
	ChangeTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
	DeleteAllTodos(c *gin.Context)
}

type todoController struct {
	tm model.TodoModel //カラム名 型  TodoModelインターフェースを定義
}

func CreateTodoContoroller(tm model.TodoModel) TodoController {
	return &todoController{tm}
}

func (tc *todoController) FetchTodos(c *gin.Context) {
	tc.tm.FetchTodos(c)

}
func (tc *todoController) AddTodo(c *gin.Context) {
	tc.tm.AddTodo(c)

}
func (tc *todoController) ChangeTodo(c *gin.Context) {
	tc.tm.ChangeTodo(c)

}
func (tc *todoController) DeleteTodo(c *gin.Context) {
	tc.tm.DeleteTodo(c)

}

func (tc *todoController) DeleteAllTodos(c *gin.Context) {
	tc.tm.DeleteAllTodos(c)

}
