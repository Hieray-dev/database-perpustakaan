package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"perpustakaan-api/config"
	"perpustakaan-api/routes"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	config.ConnectDB()

	appRoutes := routes.SetupRoutes()

	handlerWithCORS := corsMiddleware(appRoutes)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	fmt.Printf("Server running on port %s...\n", port)
	
	err := http.ListenAndServe(port, handlerWithCORS)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
