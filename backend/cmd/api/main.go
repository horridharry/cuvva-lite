package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(HealthResponse{
			Status: "ok",
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Coverly API running on port %s", port)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
