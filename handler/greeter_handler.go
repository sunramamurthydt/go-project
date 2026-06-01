package handler

import (
	"encoding/json"
	"go-project/greeter"
	"net/http"
)

type GreetResponse struct {
	Style   string `json:"style"`
	Message string `json:"message"`
}

// GreeterHandler holds a dependency on the interface, not any concrete type.
// This is DEPENDENCY INJECTION — the handler doesn't care which greeter it gets.
type GreeterHandler struct {
	greeterFactory func(style string) greeter.GreeterService
}

// NewGreeterHandler is the constructor for the handler
func NewGreeterHandler() *GreeterHandler {
	return &GreeterHandler{
		greeterFactory: greeter.NewGreeter,
	}
}

// Greet handles GET /greet?name=Sundar&style=formal
func (h *GreeterHandler) Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing 'name' query parameter", http.StatusBadRequest)
		return
	}

	style := r.URL.Query().Get("style") // optional, defaults to casual

	// Factory returns the right implementation — handler doesn't know which one
	svc := h.greeterFactory(style)

	resp := GreetResponse{
		Style:   svc.StyleName(),
		Message: svc.Greet(name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
