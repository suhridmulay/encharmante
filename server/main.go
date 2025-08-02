package main

import (
	"encharmante/server/controllers"
	"fmt"
	"net/http"
)

func genericHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/roll", controllers.HandleDiceRoll)
	http.HandleFunc("/", genericHandler)
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
