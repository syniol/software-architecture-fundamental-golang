package restful

import (
	"net/http"

	"github.com/syniol/software-architecture-fundamental-golang/pkg"
)

func NewRESTfulHealthEndpoint() (path string, handler pkg.EndpointHandler) {
	return "/health", pkg.NewRESTfulEndpointHandler(
		func(wr http.ResponseWriter, rq *http.Request) {
			_, _ = wr.Write([]byte(`{"status": "healthy"}`))
		},
		http.MethodGet,
	)
}
