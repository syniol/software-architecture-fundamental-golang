package restful

import (
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/pkg/server"
)

func NewRESTfulHealthEndpoint() (path string, handler server.EndpointHandler) {
	return "/health", func(wr http.ResponseWriter, rq *http.Request) {
		if rq.Method != http.MethodGet {
			wr.WriteHeader(http.StatusNotFound)

			return
		}

		wr.Header().Set("Content-Type", "application/json")
		wr.Write([]byte(`{"status": "healthy"}`))
	}
}
