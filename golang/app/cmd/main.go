package main

import (
	"example/todoapp/controller"
	"example/todoapp/model"
	"fmt"

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
	// todoModel(インターフェース)のインスタンスを作成
	var tm = model.CreateTodoModel()

	// todoControllerのインスタンスを作成、上で作ったtodoModelを入れる
	var tc = controller.CreateTodoContoroller(tm)

	var ro = controller.CreateRouter(tc)

	router := ro.SetupRouter()

	router.Run(":8080")
}
