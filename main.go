package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET")
	fmt.Println("GO web server is running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World!</h1>")
}
