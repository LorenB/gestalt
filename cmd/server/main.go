package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Setup logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("POST /predict", handlePredict(logger))
	mux.HandleFunc("GET /health", handleHealth(logger))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Starting server", "port", port)
	http.ListenAndServe(":"+port, mux)
}

func handlePredict(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TOOO: implement
		w.WriteHeader(http.StatusOK)
	}
}

func handleHealth(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
