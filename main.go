package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Das hier ist Ben's server. Alle request brauchen ein Passwort!")	
}

func main(){

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println("Error beim starten: ",err)
		return
	}
}

