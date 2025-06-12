package servertodolist

import (
	"net/http"
)


type Todo struct {
	Contents   string `json:"ToDoContents"`
	ExpireDate string `json:"ExpireDate"`
}

const fileName = "todos.json"

func Start() {
	http.HandleFunc("/todo/addtodo", addToDo)
	//http.HandleFunc("/todo/gettodo", handler http.Handler)
	//http.HandleFunc("/todo/removetodo", handler http.Handler)
}
