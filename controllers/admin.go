package controllers

import (
	"encoding/json"
	"net/http"
	"perpustakaan-api/config"
)

type UpdateRoleRequest struct {
	IDUser int `json:"id_user"`
	IDRole int `json:"id_role"`
}

func UpdateUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	var req UpdateRoleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	query := "UPDATE user SET id_role = ? WHERE id_user = ?"
	result, err := config.DB.Exec(query, req.IDRole, req.IDUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "User tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Role user berhasil diperbarui!",
	})
}
