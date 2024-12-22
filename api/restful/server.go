package restful

import (
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/api/restful/student/card"
	"github.com/syniol/software-architecture-fundamental-golang/pkg/restful"
)

func NewRESTfulServer() http.Handler {
	server := restful.NewServer()

	server.HandleFunc(NewRESTfulHealthEndpoint())
	server.HandleFunc(card.NewRESTfulCreateStudentCardEndpoint())
	server.HandleFunc(card.NewRESTfulStudentCardPhotoUploadEndpoint())

	return server.Handler()
}
