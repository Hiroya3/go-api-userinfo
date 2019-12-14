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
		Addr: "localhost:8080",
	}
	http.HandleFunc("/user_info/", handleRequest)
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

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var info UserInfo
	json.Unmarshal(body, &info)
	err = info.createInfo()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	info, err := getOne(id)
	if err != nil {
		return
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	//取得した構造体UserInfoを上書き
	json.Unmarshal(body, &info)
	err = info.updateInfo()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	info, err := getOne(id)
	if err != nil {
		return
	}
	err = info.deleteInfo()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
