package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"perpustakaan-api/config"
	"perpustakaan-api/routes"
)

func main() {
	config.ConnectDB()

	appRoutes := routes.SetupRoutes()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	fmt.Printf("Server running on port %s...\n", port)
	err := http.ListenAndServe(port, appRoutes)
	if err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
