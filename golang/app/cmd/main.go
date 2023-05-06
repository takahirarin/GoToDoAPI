package main

import (
	"example/todoapp/controller"
	"example/todoapp/model"
	"fmt"
	"net/http"

	// go.modで初めにモジュール名として定義したもの/インポートしたいフォルダ名？
	"flag"
)

// todoModel(インターフェース)のインスタンスを作成
var tm = model.CreateTodoModel()

// todoControllerのインスタンスを作成、上で作ったtodoModelを入れる
var tc = controller.CreateTodoContoroller(tm)

var ro = controller.CreateRouter(tc)

func migrate() {

	req := &model.Todo{
		Name:   "お風呂",
		Status: "作業中",
	}

	model.Db.Create(&req)

	fmt.Println("Migration is success!")

}

func main() {
	f := flag.String("option", "", "migrate database or not") // コマンドライン引数を取得
	flag.Parse()
	// go run cmd/main.go -option=migrateが実行された場合のみ、migrate関数を実行する。
	if *f == "migrate" {
		migrate()
	}

	http.HandleFunc("/fetch-todos", ro.FetchTodos)
	http.HandleFunc("/add-todo", ro.AddTodo)
	http.HandleFunc("/delete-todo", ro.DeleteTodo)
	http.HandleFunc("/change-todo", ro.ChangeTodo)
	http.HandleFunc("/delete-all-todo", ro.DeleteAllTodos)
	http.ListenAndServe(":8080", nil)

}
