package main

import(
	"database/sql"
	_ "github.com/lib/qp"
)

var Db *sql.Db

func init(){
	var err error
	DB, err := sql.Open("postgres","user=dbuser dbname=user_info password=pass sslmode=disable")
	if err != nil{
		panic(err)
	}
}
