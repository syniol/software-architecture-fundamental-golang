package server

import "net/http"

type EndpointHandler = func(http.ResponseWriter, *http.Request)
