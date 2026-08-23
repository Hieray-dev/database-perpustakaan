package main

import (
	"net/http"
	"perpustakaan-api/config"
	"perpustakaan-api/controllers"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/login", controllers.LoginHandler)

	http.HandleFunc("/peminjaman/riwayat", controllers.GetRiwayatHandler)

	http.ListenAndServe(":8080", nil)
}
