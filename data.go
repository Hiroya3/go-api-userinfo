package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//データベースへのハンドル（構造体）
var Db *sql.Db

func init() {
	var err error
	//Openを使っているが、これはsql.Dbに接続に必要な構造体を設定するのみ
	DB, err := sql.Open("postgres", "user=dbuser dbname=user_info password=pass sslmode=disable")
	if err != nil {
		panic(err)
	}
}
