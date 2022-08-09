package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "码神之路开启，Go Blog"
	indexData.Desc = "盘古开天辟地"
	bytes, _ := json.Marshal(indexData)
	w.Write(bytes)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "码神之路开启，Go Blog"
	indexData.Desc = "盘古开天辟地"
	t := template.New("index.html")
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
