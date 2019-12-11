package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//データベースへのハンドル（構造体）
var Db *sql.DB

func init() {
	var err error
	//Openを使っているが、これはsql.Dbに接続に必要な構造体を設定するのみ
	DB, err := sql.Open("postgres", "user=dbuser dbname=user_info password=pass sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func getOne(id int) (userInfo UserInfo, err error) {
	userInfo = UserInfo{}
	err = Db.QueryRow("select id, name, age from user_info where id = $1", id).Scan(&userInfo.Id, &userInfo.Name, &userInfo.Age)
	return
}
