package main

import (
	"fmt"
	"net/http"
	"server/servertodolist"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Das hier ist Ben's server. Alle request brauchen ein Passwort!")
	ascii := `   
   .--.
  |o_o |
  |:_/ |
 //   \ \
(|     | )
/'\_   _/` + "`" + `\
\___)=(___/`

	fmt.Fprint(w, ascii)
}

func main() {
	http.HandleFunc("/", handler)
	servertodolist.Start()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error beim starten: ", err)
		return
	}
}
