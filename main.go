package main

import (
	"fmt"
	"net/http"
	"os"
	"perpustakaan-api/config"
	"perpustakaan-api/routes"
)

func main() {
	config.ConnectDB()

	appRoutes := routes.SetupRoutes()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	fmt.Println("Server running on port :8080...")
	err := http.ListenAndServe(":8080", appRoutes)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
