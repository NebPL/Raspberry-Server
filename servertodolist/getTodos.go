package servertodolist

import (
	"encoding/json"
	"os"
)


type Todo struct{
	Contents string `json:"ToDoContents"`
	ExpireDate string `json:"ExpireDate"`
}

const fileName = "todos.json"

func getTodos()(Todo, error){
  var todo Todo

	 file, err := os.Open(fileName)
    if err != nil {
        if os.IsNotExist(err) {
            return Todo{}, nil // Datei existiert noch nicht, kein Problem
        }
        return todo, err
    }
    defer file.Close()

    err = json.NewDecoder(file).Decode(&todo)
    return todo, err
	}
