package main

import (
	"fmt"
	"log"
	"net/http"
	"server/servertodolist"
)

const AuthPasswort = "uoyCjXEPJVWnpaedHiO"

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

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := r.Header.Get("Passwort")

		if password == "" || password != AuthPasswort {
			log.Println("Jemand versuchte sich einzulogen mit Falschen passwort!")

			http.Error(w, "Unauthorized: Invalid or missing password", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var test int

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	servertodolist.Start(mux)

	checkedMux := authMiddleware(mux)

	err := http.ListenAndServe(":8080", checkedMux)
	if err != nil {
		fmt.Println("Error beim starten: ", err)
		return
	}
}
