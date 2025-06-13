package servertodolist

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getTodosApi(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		fmt.Fprintln(w, "Only Post not ", r.Method)
		return
	}

	var todos []Todo // Deklariere einen Slice von Todo-Structs

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(w, "Error: ", err)
		return
	}
	defer file.Close()

	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Get Todos")
	fmt.Println("")
	log.Printf("Debug: New Debug Request")

	// Versuche, den Inhalt der Datei in den Slice von Todos zu decodieren
	err = json.NewDecoder(file).Decode(&todos)
	if err != nil {
		fmt.Errorf("Fehler beim Decodieren der Todos: %w", err) // Füge fmt für Fehlerformatierung hinzu
	}

	//log.Printf("Debug: %d Todo length", len(todos))

	for _, content := range todos {
		log.Println("Contents: ", content.Contents)
		log.Println("ExpireDate in Dayes: ", content.ExpireTime)

		fmt.Println(" ")
	}

	log.Printf("Debug: %d Todo length", len(todos))
	fmt.Println(" ")

	jsonData, err := json.Marshal(todos)

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
