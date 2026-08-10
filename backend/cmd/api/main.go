package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/horridharry/cuvva-lite/backend/internal/database"
	"github.com/horridharry/cuvva-lite/backend/internal/quote"
	"github.com/horridharry/cuvva-lite/backend/internal/vehicle"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	_ = godotenv.Load()

	ctx := context.Background()

	db, databaseError := database.NewPostgresPool(
		ctx,
		os.Getenv("DATABASE_URL"),
	)

	if databaseError != nil {
		log.Fatal(databaseError)
	}
	defer db.Close()

	vehicleRepository := vehicle.NewRepository(db)
	quoteRepository := quote.NewRepository(db)

	quoteService := quote.NewService(
		quoteRepository,
		vehicleRepository,
	)

	quoteHandler := quote.NewHandler(quoteService)
	vehicleHandler := vehicle.NewHandler(vehicleRepository)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if responseError := json.NewEncoder(w).Encode(HealthResponse{
			Status: "ok",
		}); responseError != nil {
			http.Error(
				w,
				"Failed to encode reponse",
				http.StatusInternalServerError,
			)
		}
	})

	mux.HandleFunc(
		"POST /api/quotes",
		quoteHandler.Create,
	)

	mux.HandleFunc(
		"GET /api/vehicles/{registration}",
		vehicleHandler.GetByRegistration,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Api running on port %s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
