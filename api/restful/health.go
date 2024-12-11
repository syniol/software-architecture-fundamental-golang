package restful

import (
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/pkg/server"
)

func NewRESTfulHealthEndpoint() (path string, handler server.EndpointHandler) {
	return "/health", server.NewRESTfulEndpointHandler(
		func(wr http.ResponseWriter, rq *http.Request) {
			_, _ = wr.Write([]byte(`{"status": "healthy"}`))
		},
		http.MethodGet,
	)
}
