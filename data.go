package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB //データベースへのハンドル（構造体）

func init() {
	var err error
	//Openを使っているが、これはsql.Dbに接続に必要な構造体を設定するのみ
	db, err = sql.Open("postgres", "user=dbuser dbname=user_info password=pass sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func getOne(id int) (userInfo UserInfo, err error) {
	userInfo = UserInfo{}
	err = db.QueryRow("select id, name, age from user_info where id = $1", id).Scan(&userInfo.Id, &userInfo.Name, &userInfo.Age)
	return
}

func (info *UserInfo) createInfo() (err error) {
	err = db.QueryRow("insert into user_info (name,age) values ($1,$2) returning id", info.Name, info.Age).Scan(&info.Id)
	return
}

func (info *UserInfo) updateInfo() (err error) {
	//updateで返されるメソッドは関係ない
	_, err = db.Exec("update user_info set name = $2, age = $3 where id = $1", info.Id, info.Name, info.Age)
	return
}

func (info *UserInfo) deleteInfo() (err error) {
	_, err = db.Exec("delete from user_info where id = $1", info.Id)
	return
}
