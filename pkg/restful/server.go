package restful

import "net/http"

type Server struct {
	server *http.ServeMux
}

func (s *Server) Handler() http.Handler {
	return s.server
}

func (s *Server) HandleFunc(pattern string, handler EndpointHandler) {
	s.HandleFunc(pattern, handler)
}

func NewServer() *Server {
	return &Server{
		server: http.NewServeMux(),
	}
}
