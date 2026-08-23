package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan-api/config"
	"perpustakaan-api/models"
)

func GetRiwayatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	query := `
	 SELECT
	  p.id_peminjaman,
	  u.username,
	  u.is_active,
	  p.tanggal_peminjaman
	 FROM peminjaman p
	 JOIN user u ON p.id_petugas = u.id_user`

	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listRiwayat []models.Riwayat
	for rows.Next() {
		var rwt models.Riwayat
		var isActive int
		err := rows.Scan(&rwt.IDPeminjaman, &rwt.Pustakawan, &isActive, &rwt.TanggalPinjam)
		if err != nil {
			continue
		}

		if isActive == 1 {
			rwt.StatusPetugas = "Aktif"
		} else {
			rwt.StatusPetugas = "Nonaktif (Resign)"
		}

		listRiwayat = append(listRiwayat, rwt)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listRiwayat)
}
