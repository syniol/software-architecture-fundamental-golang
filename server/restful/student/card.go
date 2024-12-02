package student

import "net/http"

func NewRESTfulStudentCardEndpoint(server *http.ServeMux) {
	server.HandleFunc("/v1/student/card", func(wr http.ResponseWriter, rq *http.Request) {
		if rq.Method == http.MethodPost {
			wr.WriteHeader(http.StatusNotFound)

			return
		}

		wr.Header().Set("Content-Type", "application/json")
		wr.Write([]byte(`{"status": "to be implemented"}`))
	})
}
