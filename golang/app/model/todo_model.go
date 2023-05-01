package model

import (
	// "database/sql"
	// "fmt"
	// "github.com/oklog/ulid"
	// "math/rand"
	"net/http"
	
	// "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type TodoModel interface {
	FetchTodos() ([]*Todo, error) //メソッド名 (引数) (リターン)
	AddTodo(r *http.Request) (*Todo, error)
	ChangeTodo(r *http.Request) (*Todo, error)
	DeleteTodo(r *http.Request) (error)
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

func (tm *todoModel) FetchTodos() (todos []*Todo, err error) {
	// 構造体todoModelに対するメソッド、構造体todoModelはインターフェースTodoModelと
	// 同じメソッドを持っているのでインターフェースTodoModel型に変換できる

	// sql := `SELECT id, name, status FROM todos`

	// rows, err := Db.Query(sql)
	err = Db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	
	return todos, nil

}

func (tm *todoModel) AddTodo(r *http.Request) (*Todo, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, nil
	}

	req := Todo{
		Name:   r.FormValue("name"),
		Status: r.FormValue("status"),
	}

	result := Db.Create(&req)
	if result.Error != nil {
		return &req, result.Error
	}

	return &req, nil

}

func (tm *todoModel) ChangeTodo(r *http.Request) (*Todo, error) {
	err := r.ParseForm() // Responseからvalueを取り出すための準備

	if err != nil {
		return nil, nil
	}

	todo := Todo{}
	Db.First(&todo, r.FormValue("id"))
	if todo.Status == "作業中"{
		todo.Status = "完了"
	}

	result := Db.Save(&todo)

	if result.Error != nil {
		return &todo, result.Error
	}

	return &todo, nil
}

func (tm *todoModel) DeleteTodo(r *http.Request) (error) {
	err := r.ParseForm()

	if err != nil {
		return nil
	}

	result := Db.Delete(&Todo{}, r.FormValue("id"))

	if result.Error != nil {
		return result.Error
	}

	return nil
}
