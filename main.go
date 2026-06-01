package main

import (
	"encoding/json"
	"fmt"
	"go-project/handler"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing 'name' query parameter", http.StatusBadRequest)
		return
	}

	resp := Response{
		Message: fmt.Sprintf("Hey %s! Welcome aboard — glad to have you here!", name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	greeterHandler := handler.NewGreeterHandler()

	// Module 1 route — kept from before
	http.HandleFunc("/welcome", welcomeHandler)

	// Module 2 route — OOP via interface + factory
	http.HandleFunc("/greet", greeterHandler.Greet)

	fmt.Println("Server running at http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
