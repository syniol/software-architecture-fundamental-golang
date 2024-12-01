package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", func(wr http.ResponseWriter, rq *http.Request) {
		wr.Header().Set("Content-Type", "application/json")
		wr.Write([]byte(`{"status": "healthy"}`))
	})

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}
