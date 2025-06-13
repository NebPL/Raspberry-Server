package servertodolist

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func addToDo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("Add Todos")
	fmt.Println("")

	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Only POST not ", r.Method)
		log.Println("Only POST not ", r.Method)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad Request: Could not parse todo from JSON: %v", err), http.StatusBadRequest)
		log.Printf("DEBUG: addToDo: Fehler beim Parsen des Todo-JSON: %v, Body: %s", err, string(body))
		return
	}

	var newTodo Todo

	err = json.Unmarshal(body, &newTodo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad Request: Could not parse todo from JSON: %v", err), http.StatusBadRequest)
		log.Printf("DEBUG: addToDo: Fehler beim Parsen des Todo-JSON: %v, Body: %s", err, string(body))
		return
	}

	todos, err := readTodosFromFile()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error whith Reading Fiele: %v", err), http.StatusBadRequest)
		log.Printf("DEBUG: addToDo: Fehler beim Parsen des Todo-JSON: %v, Body: %s", err, string(body))
		return
	}

	todos = append(todos, newTodo)
	log.Printf("DEBUG: addToDo: Neues ToDo hinzugefügt. Gesamtzahl: %d", len(todos))

	err = writeTodosToFile(todos)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal Server Error: Could not save todos: %v", err), http.StatusInternalServerError)
		log.Printf("DEBUG: addToDo: Fehler beim Speichern der Todos: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201 Created ist passend für das Hinzufügen einer Ressource
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "ToDo added successfully"})
	log.Println("DEBUG: addToDo: ToDo erfolgreich hinzugefügt und Antwort gesendet.")
	fmt.Println("")
}

func readTodosFromFile() ([]Todo, error) {
	var todos []Todo

	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil // Datei existiert nicht, gib leeres Array zurück
		}
		return nil, fmt.Errorf("Fehler beim Öffnen der Todo-Datei: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&todos)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim Decodieren der Todos aus Datei: %w", err)
	}
	return todos, nil
}

// writeTodosToFile ist eine Hilfsfunktion, um ToDos in die Datei zu schreiben
func writeTodosToFile(todos []Todo) error {
	jsonData, err := json.MarshalIndent(todos, "", "  ") // MarshalIndent für schönere Formatierung
	if err != nil {
		return fmt.Errorf("Fehler beim Marshaling der Todos: %w", err)
	}

	err = ioutil.WriteFile(fileName, jsonData, 0644) // 0644 sind Dateiberechtigungen (rw-r--r--)
	if err != nil {
		return fmt.Errorf("Fehler beim Schreiben der Todos in Datei: %w", err)
	}
	return nil
}
