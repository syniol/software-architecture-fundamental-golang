package restful

import (
	"github.com/syniol/software-architecture-fundamental-golang/pkg/restful"
	"net/http"
)

func NewRESTfulHealthEndpoint() (path string, handler restful.EndpointHandler) {
	return "/health", restful.NewRESTfulEndpointHandler(
		func(wr http.ResponseWriter, rq *http.Request) {
			_, _ = wr.Write([]byte(`{"status": "healthy"}`))
		},
		http.MethodGet,
	)
}
