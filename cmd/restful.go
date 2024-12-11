package main

import (
	"log"
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/api/restful"
)

func main() {
	err := http.ListenAndServe(":8080", restful.NewRESTfulServer())
	if err != nil {
		log.Fatal(err)
	}
}
