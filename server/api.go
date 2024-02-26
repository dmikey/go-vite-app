package main

import (
	"encoding/json"
	"net/http"

	go_vite_app "github.com/dmikey/go-vite-app/server/proto"
)

// apiHandler is an example API handler function.
func apiHandler(w http.ResponseWriter, r *http.Request) {
	app := &go_vite_app.App{
		Name: "hello-world",
	}
	
	// response := map[string]string{"message": "Hello from the API"}
	json.NewEncoder(w).Encode(app)
}

// RegisterAPIRoutes sets up the API routes.
func RegisterAPIRoutes() {
	http.HandleFunc("/api", apiHandler)
	// Add more API routes here.
}
