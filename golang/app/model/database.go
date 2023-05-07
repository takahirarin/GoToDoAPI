package model

import (
	"fmt"
	"log"

	// "os"
	// "database/sql"
	_ "github.com/lib/pq"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	_ "github.com/microsoft/go-mssqldb"
)

// データベースへのハンドル作成
// var Db *sql.DB
var Db *gorm.DB
var server = "todoapi-database.database.windows.net"
var port = 1433
var user = "takahirarin"
var password = "Piyorinko_0118"
var database = "ToDoAPIDataBase"

// データベースへ接続
func init() {
	var err error
	//dsn := "host=db user=rin password=rin dbname=rin sslmode=disable"

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	Db, err = gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	//Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&Todo{})

	fmt.Println("Connection has been established!")
}
