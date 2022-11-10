package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)

	http.HandleFunc("/tickets", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			getTickets(w, r)

		case http.MethodPost:
			addTicket(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

	})
}
