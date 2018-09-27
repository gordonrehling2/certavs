package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s endpoint\n", r.URL.Path[1:])
	state := struct {
		Status string    `json:"status"`
		Errors [0]string `json:"errors"`
	}{Status: "OK"}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(state); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

