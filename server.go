package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type UserInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/post", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	//path.Baaeでid取得
	//strconv.Atoiで文字列のidを整数に変換
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	info, err := getOne(id)
	return_info, err := json.MarshalIndent(&info, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(return_info)
	return
}
