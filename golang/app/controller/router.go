package controller

import (
	"os"

	"github.com/gin-gonic/gin"
)

type router struct {
	tc TodoController
}

func CreateRouter(tc TodoController) *router {
	return &router{tc}
}

func (r *router) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		// プリフライトリクエスト用に設定している。
		c.Header("Access-Control-Allow-Headers", "*")
		// CORSエラー対策。APIを叩くフロント側のURLを渡す。
		c.Header("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
		// 返却する値のContent-Typeを設定。
		c.Header("Content-Type", "application/json")
		c.Next()
	})

	router.GET("/fetch-todos", r.tc.FetchTodos)
	router.POST("/add-todo", r.tc.AddTodo)
	router.DELETE("/delete-todo/:id", r.tc.DeleteTodo)
	router.PUT("/change-todo/:id", r.tc.ChangeTodo)
	router.GET("/delete-all-todo", r.tc.DeleteAllTodos)

	return router
}
