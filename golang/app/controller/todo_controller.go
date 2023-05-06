package controller

import (
	"encoding/json"
	"example/todoapp/model"
	"fmt"
	"net/http"
)

type TodoController interface {
	FetchTodos(w http.ResponseWriter, r *http.Request)
	AddTodo(w http.ResponseWriter, r *http.Request)
	ChangeTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
	DeleteAllTodos(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tm model.TodoModel //カラム名 型  TodoModelインターフェースを定義
}

func CreateTodoContoroller(tm model.TodoModel) TodoController {
	return &todoController{tm}
}

func (tc *todoController) FetchTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := tc.tm.FetchTodos()

	if err != nil {
		fmt.Fprint(w, err) // wに書き込み
		return
	}

	json, err := json.Marshal(todos)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))

}
func (tc *todoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	result, err := tc.tm.AddTodo(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))

}
func (tc *todoController) ChangeTodo(w http.ResponseWriter, r *http.Request) {
	result, err := tc.tm.ChangeTodo(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))

}
func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	err := tc.tm.DeleteTodo(r)

	if err != nil {
		fmt.Fprint(w, err) // wに書き込み
		return
	}

	json, err := json.Marshal(r.FormValue("id"))

	if err != nil {
		fmt.Fprint(w, err) // wに書き込み
		return
	}

	fmt.Fprint(w, string(json))

}

func (tc *todoController) DeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	result, err := tc.tm.DeleteAllTodos()

	if err != nil {
		fmt.Fprint(w, err) // wに書き込み
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err) // wに書き込み
		return
	}

	fmt.Fprint(w, string(json))

}
