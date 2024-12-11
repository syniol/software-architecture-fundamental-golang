package restful

import (
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/api/restful/student"
)

func NewRESTfulServer() http.Handler {
	server := http.NewServeMux()

	server.HandleFunc(NewRESTfulHealthEndpoint())
	server.HandleFunc(student.NewRESTfulCreateStudentCardEndpoint())

	return server
}
