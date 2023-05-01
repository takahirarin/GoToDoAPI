package model

import (
	"database/sql"
	"fmt"
	"github.com/oklog/ulid"
	"math/rand"
	"net/http"
	"time"
)

type TodoModel interface {
	FetchTodos() ([]*Todo, error) //メソッド名 (引数) (リターン)
	AddTodo(r *http.Request) (sql.Result, error)
	ChangeTodo(r *http.Request) (sql.Result, error)
	DeleteTodo(r *http.Request) (sql.Result, error)
}

type todoModel struct {
}

type Todo struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func CreateTodoModel() TodoModel { // ←戻り値の型がTodoModel(=interface)になっている
	return &todoModel{}
}

// 作成されたTodoModelはインターフェースなので具体的なメソッドの実装内容(下の実装)まで見れない

func (tm *todoModel) FetchTodos() ([]*Todo, error) {
	// 構造体todoModelに対するメソッド、構造体todoModelはインターフェースTodoModelと
	// 同じメソッドを持っているのでインターフェースTodoModel型に変換できる

	sql := `SELECT id, name, status FROM todos`

	rows, err := Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo

	for rows.Next() {
		var (
			id, name, status string
		)
		if err := rows.Scan(&id, &name, &status); err != nil {
			return nil, err
		}

		todos = append(todos, &Todo{
			Id:     id,
			Name:   name,
			Status: status,
		})
	}

	return todos, nil

}

func (tm *todoModel) AddTodo(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, nil
	}

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	req := Todo{
		Id:     id.String(),
		Name:   r.FormValue("name"),
		Status: r.FormValue("status"),
	}

	fmt.Println(r.FormValue("name"))
	fmt.Println(req.Id, req.Name, req.Status)

	sql := `INSERT INTO todos(id, name, status) VALUES($1, $2, $3)`

	result, err := Db.Exec(sql, req.Id, req.Name, req.Status)
	if err != nil {
		return result, err
	}

	return result, nil

}

func (tm *todoModel) ChangeTodo(r *http.Request) (sql.Result, error) {
	err := r.ParseForm() // Responseからvalueを取り出すための準備

	if err != nil {
		return nil, nil
	}

	sql := `UPDATE todos SET status = $1 WHERE id = $2`

	afterStatus := "作業中"
	if r.FormValue("status") == "作業中" {
		afterStatus = "完了"
	}

	result, err := Db.Exec(sql, afterStatus, r.FormValue("id"))

	if err != nil {
		return result, err
	}

	return result, nil
}

func (tm *todoModel) DeleteTodo(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, nil
	}

	sql := `DELETE FROM todos WHERE id = $1`

	result, err := Db.Exec(sql, r.FormValue("id"))

	if err != nil {
		return result, err
	}

	return result, nil
}
