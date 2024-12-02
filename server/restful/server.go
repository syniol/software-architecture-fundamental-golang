package restful

import (
	"github.com/syniol/software-architecture-fundamental-golang/server/restful/student"
	"net/http"
)

func NewServer() *http.ServeMux {
	server := http.NewServeMux()

	NewRESTfulHealthEndpoint(server)
	student.NewRESTfulStudentCardEndpoint(server)

	return server
}
