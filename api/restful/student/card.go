package student

import (
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/pkg/server"
)

func NewRESTfulStudentCardEndpoint() (path string, handler server.EndpointHandler) {
	return "/v1/student/card", func(wr http.ResponseWriter, rq *http.Request) {
		if rq.Method == http.MethodPost {
			wr.WriteHeader(http.StatusNotFound)

			return
		}

		wr.Header().Set("Content-Type", "application/json")
		wr.Write([]byte(`{"status": "to be implemented"}`))
	}
}
