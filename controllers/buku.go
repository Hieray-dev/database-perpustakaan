package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan-api/config"
	"perpustakaan-api/models"
)

func GetKatalogBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query("SELECT id_buku, judul, penulis, deskripsi, gambar, stok FROM buku")
	if err != nil {
		http.Error(w, `{"message":"Gagal mengambil data buku"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listBuku []models.Buku
	for rows.Next() {
		var b models.Buku
		rows.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Deskripsi, &b.Gambar, &b.Stok)
		listBuku = append(listBuku, b)
	}

	json.NewEncoder(w).Encode(listBuku)
}
