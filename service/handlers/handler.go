package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/gordonrehling2/certavs/service"
)

type Handler struct {
	RFE service.RfeService
}

func NewHandler(rfe service.RfeService) *Handler {
	return &Handler {
		RFE: rfe,
	}
}

func (h *Handler) RfeList() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}

func (h *Handler) RfeProductList() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}

func (h *Handler) RfeProductCreate() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}

func (h *Handler) HealthCheck() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "%s endpoint\n", r.URL.Path[1:])
		state := struct {
			Status string    `json:"status"`
			Errors [0]string `json:"errors"`
		} {Status: "OK"}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(state); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
