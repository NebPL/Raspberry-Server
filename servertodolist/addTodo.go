package servertodolist

import (
	"fmt"
	"net/http"
)

func addToDo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Only Post not ", r.Method)
	}

}
