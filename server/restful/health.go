package restful

import "net/http"

func NewRESTfulHealthEndpoint(server *http.ServeMux) {
	server.HandleFunc("/health", func(wr http.ResponseWriter, rq *http.Request) {
		if rq.Method != http.MethodGet {
			wr.WriteHeader(http.StatusNotFound)

			return
		}

		wr.Header().Set("Content-Type", "application/json")
		wr.Write([]byte(`{"status": "healthy"}`))
	})
}
