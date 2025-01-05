package httpserver

import "net/http"

type HttpServer struct {
	Address string
}

func NewHttpServer(address string) *HttpServer {
	return &HttpServer{
		Address: address,
	}
}

func (server *HttpServer) Start() {
	http.ListenAndServe(server.Address, nil)
}