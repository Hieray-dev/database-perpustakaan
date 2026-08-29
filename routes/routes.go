package routes

import (
	"net/http"
	"perpustakaan-api/controllers"
	"perpustakaan-api/middleware"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", controllers.RegisterHandler)
	mux.HandleFunc("/login", controllers.LoginHandler)
	mux.HandleFunc("/buku", controllers.GetKatalogBuku)

	mux.Handle("/booking", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateBooking)))
	mux.Handle("/histori/riwayat", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetRiwayatHandler)))

	mux.Handle("/peminjaman/confirm", middleware.AuthMiddleware(http.HandlerFunc(controllers.ConfirmPinjam)))

  mux.Handle("/petugas/create", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreatePetugas)))

	return mux
}
