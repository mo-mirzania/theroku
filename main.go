package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(determineListenAddress(), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! from Heroku!")
}

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":9090"
	}
	return ":" + port
}
