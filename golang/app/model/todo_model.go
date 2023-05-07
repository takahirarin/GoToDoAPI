package model

import (
	// "database/sql"

	// "github.com/oklog/ulid"
	// "math/rand"

	"net/http"

	// "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type TodoModel interface {
	FetchTodos(c *gin.Context) //メソッド名 (引数) (リターン)
	AddTodo(c *gin.Context)
	ChangeTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
	DeleteAllTodos(c *gin.Context)
}

type todoModel struct {
}

type Todo struct {
	gorm.Model
	Name   string
	Status string
}

func CreateTodoModel() TodoModel { // ←戻り値の型がTodoModel(=interface)になっている
	return &todoModel{}
}

// 作成されたTodoModelはインターフェースなので具体的なメソッドの実装内容(下の実装)まで見れない

func (tm *todoModel) FetchTodos(c *gin.Context) {
	var todos []*Todo
	if err := Db.Find(&todos).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error()) //errじゃなくてerr.Error()??
		return
	}

	c.JSON(http.StatusOK, todos)

}

func (tm *todoModel) AddTodo(c *gin.Context) {
	var req Todo
	if err := c.ShouldBindJSON(&req); err != nil { // Todo型の変数reqにgin.Contextで送られてきたjsonデータを入れ込む
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := Db.Create(&req).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, req)

}

func (tm *todoModel) ChangeTodo(c *gin.Context) {
	var todo Todo
	id := c.Param("id")
	// body, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println("Request Body:", string(body))

	if err := Db.First(&todo, id).Error; err != nil {
		c.String(http.StatusNotFound, "Todo not found")
		return
	}

	if todo.Status == "作業中" {
		todo.Status = "完了"
	}

	if err := Db.Save(&todo).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)

}

func (tm *todoModel) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if err := Db.Delete(&Todo{}, id).Error; err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
	// gin.H -> map型のデータ構造。c.JSONは第一引数にhttpステータスコード、第二引数にレスポンスの本文をgin.H型で返す
}

func (tm *todoModel) DeleteAllTodos(c *gin.Context) {
	var todos []*Todo

	if err := Db.Find(&todos).Error; err != nil {
		return
	}

	if len(todos) == 0 {
		return
	}

	// 全レコードを削除する
	if err := Db.Delete(&todos).Error; err != nil {
		return
	}

	c.JSON(http.StatusOK, todos)
}
