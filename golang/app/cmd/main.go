package main

import (
	"fmt"
	"net/http"
	"example/todoapp/model"
	"example/todoapp/controller"
	// go.modで初めにモジュール名として定義したもの/インポートしたいフォルダ名？
	"flag"
)

//todoModel(インターフェース)のインスタンスを作成
var tm = model.CreateTodoModel()
//todoControllerのインスタンスを作成、上で作ったtodoModelを入れる
var tc = controller.CreateTodoContoroller(tm)

var ro = controller.CreateRouter(tc)

func migrate(){
	sql := `INSERT INTO todos(id, name, status) VALUES('00000000000000000000000000','買い物', '作業中'),('00000000000000000000000001','洗濯', '作業中'),('00000000000000000000000002','皿洗い', '完了');`
	_, err := model.Db.Exec(sql)// database.goで定義したDb
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Migration is success!")

}

func main(){
	f := flag.String("option", "", "migrate database or not")// コマンドライン引数を取得
	flag.Parse()
    // go run cmd/main.go -option=migrateが実行された場合のみ、migrate関数を実行する。
	if *f == "migrate" {
		migrate()
	}

	http.HandleFunc("/fetch-todos", ro.FetchTodos)
	http.HandleFunc("/add-todo", ro.AddTodo)
	http.HandleFunc("/delete-todo", ro.DeleteTodo)
	http.HandleFunc("/change-todo", ro.ChangeTodo)
	http.ListenAndServe(":8080", nil)

}