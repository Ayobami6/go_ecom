package root

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (r *RootHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", r.handleRoot).Methods("GET")
}

func (r *RootHandler) handleRoot(w http.ResponseWriter, _ *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to Go API!"})
}