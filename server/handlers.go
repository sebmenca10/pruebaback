package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

func getTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func addTicket(w http.ResponseWriter, r *http.Request) {

	ticket := &Ticket{}

	err := json.NewDecoder(r.Body).Decode()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "%v", err)
		return
	}

	tickets = append(tickets, ticket)
	fmt.Fprintf(w, "Ticket Creado")
}
