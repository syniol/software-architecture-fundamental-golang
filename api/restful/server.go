package restful

import (
	"github.com/syniol/software-architecture-fundamental-golang/api/restful/student/card"
	"net/http"
)

func NewRESTfulServer() http.Handler {
	server := http.NewServeMux()

	server.HandleFunc(NewRESTfulHealthEndpoint())
	server.HandleFunc(card.NewRESTfulCreateStudentCardEndpoint())
	server.HandleFunc(card.NewRESTfulStudentCardPhotoUploadEndpoint())

	return server
}
