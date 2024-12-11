package server

import (
	"net/http"
	"strings"
)

type EndpointHandler = func(http.ResponseWriter, *http.Request)

func NewRESTfulEndpointHandler(handler EndpointHandler, allowedMethod string) EndpointHandler {
	return func(wr http.ResponseWriter, rq *http.Request) {
		if rq.Method != allowedMethod {
			wr.WriteHeader(http.StatusNotFound)

			return
		}

		wr.Header().Set("Content-Type", "application/json")
		handler(wr, rq)
	}
}

func NewRESTfulEndpointAuthenticatedHandler(handler EndpointHandler, allowedMethod string) EndpointHandler {
	return func(wr http.ResponseWriter, rq *http.Request) {
		if rq.Method != allowedMethod {
			wr.WriteHeader(http.StatusNotFound)

			return
		}

		token := rq.Header.Get("Authorization")
		tokenSplit := strings.SplitAfter(token, "Bearer")
		if len(tokenSplit) <= 1 {
			wr.WriteHeader(http.StatusForbidden)

			return
		}

		wr.Header().Set("Content-Type", "application/json")
		handler(wr, rq)
	}
}
