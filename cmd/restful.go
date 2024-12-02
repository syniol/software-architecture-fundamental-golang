package main

import (
	"log"
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/server/restful"
)

func main() {
	err := http.ListenAndServe(":8080", restful.NewServer())
	if err != nil {
		log.Fatal(err)
	}
}
