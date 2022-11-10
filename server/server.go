package server

import "net/http"

type Ticket struct {
	id int
	usuario string
	fcreacion string
	factualizacion string
	estatus string
}

var tickets = []*Ticket=[]


func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
