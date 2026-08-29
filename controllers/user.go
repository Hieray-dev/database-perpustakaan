package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan-api/config"
	"perpustakaan-api/models"
)

func CreatePetugas(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"message":"Format JSON tidak valid"}`, http.StatusBadRequest)
		return
	}

	query := `INSERT INTO user (username, password, id_role, id_shift, is_active) VALUES (?, ?, ?, ?, 1)`
	_, err = config.DB.Exec(query, req.Username, req.Password, req.IDRole, req.IDShift)
	if err != nil {
		http.Error(w, `{"message":"Gagal menambah petugas: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Petugas berhasil ditambahkan",
	})
}
