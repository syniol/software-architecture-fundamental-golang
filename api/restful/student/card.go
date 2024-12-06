package student

import (
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/internal/student/card"
	"github.com/syniol/software-architecture-fundamental-golang/pkg/server"
)

func NewRESTfulCreateStudentCardEndpoint() (path string, handler server.EndpointHandler) {
	return "/v1/student/card", func(wr http.ResponseWriter, rq *http.Request) {
		createStudentCardEndpointHandler(card.NewStudentCardRepository(), wr, rq)
	}
}
