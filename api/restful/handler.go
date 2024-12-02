package restful

import (
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/api/restful/student"
)

func NewRESTfulHandler() http.Handler {
	server := http.NewServeMux()

	server.HandleFunc(NewRESTfulHealthEndpoint())
	server.HandleFunc(student.NewRESTfulStudentCardEndpoint())

	return server
}
