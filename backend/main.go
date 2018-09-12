package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Record struct {
	Source  string
	Target  string
	Amount  int
	Concept string
}

var records = []Record{
	Record{Source: "Employer", Target: "Bank", Amount: 100000, Concept: "Salary"},
	Record{Source: "Bank", Target: "Amazon", Amount: 2500, Concept: "Book"},
	Record{Source: "Bank", Target: "Aldi", Amount: 6200, Concept: "Groceries"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/records", recordsHandler)
	http.ListenAndServe(":8000", handlers.CORS()(router))
}

func recordsHandler(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(records)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
