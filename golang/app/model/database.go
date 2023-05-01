package model

import (
	"fmt"
	// "os"
	// "database/sql"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
    "gorm.io/gorm"
	
)

// データベースへのハンドル作成
// var Db *sql.DB
var Db *gorm.DB

// データベースへ接続
func init() {
	var err error
	//dsn := fmt.Sprintf("%s:%s@tcp(db:5432)/%s?charset=utf8", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))
	dsn := "host=db user=rin password=rin dbname=rin sslmode=disable"
	// Db, err = sql.Open("postgres", dsn)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// sql := `CREATE TABLE IF NOT EXISTS todos(
	// 	id varchar(26) not null,
	// 	name varchar(100) not null,
	// 	status varchar(100) not null
	// )`

	Db.AutoMigrate(&Todo{})

	// _, err = Db.Exec(sql)

	fmt.Println("Connection has been established!")
}
