package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StartWebServer() error {
	fmt.Println("Start Web Server!")
	http.HandleFunc("/api/todos", getTodos)
	return http.ListenAndServe(":8080", nil)
}

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	todo1 := Todo{
		Id:        1,
		Title:     "チャーハン",
		Completed: true,
	}

	todo2 := Todo{
		Id:        2,
		Title:     "豚肉も入れるよ！",
		Completed: false,
	}

	todos := []Todo{todo1, todo2}

	responseBody, err := json.Marshal(todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responseBody)
}