package servertodolist

import (
	"net/http"
)

type Todo struct {
	Contents   string `json:"ToDoContents"`
	ExpireTime int    `json:"ExpireTime"`
}

const fileName = "/Users/ben/home/programming/personal/RaspberryServer/servertodolist/todos.json"

func Start(mux *http.ServeMux) {
	mux.HandleFunc("/todo/addtodo", addToDo)
	mux.HandleFunc("/todo/gettodo", getTodosApi)
	//http.HandleFunc("/todo/removetodo", handler http.Handler)
}
