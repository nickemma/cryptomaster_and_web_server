package main

import (
	"fmt"
	"net/http"

	"github.com/nickemma/cryptomaster_and_web_server/apis"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/api/hello", getData)

	server.HandleFunc("/api/data", apis.HandleGet)
	server.HandleFunc("/api/data/", apis.HandleGetById)
	server.HandleFunc("/api/data/new", apis.HandlePost)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Error in opening the server")
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Full Time Go Dev!"))
}
