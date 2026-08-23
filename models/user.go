package models

type User struct {
	ID       int    `json:"id_user"`
	Username string `json:"username"`
	Password string `json:"password"`
	IDRole   int    `json:"id_role"`
	IDShift  *int   `json:"id_shift"`
	IsActive int    `json:"is_active"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IDRole   int    `json:"id_role"`
	IDShift  int    `json:"id_shift"`
}

type LoginShiftResponse struct {
	Message         string `json:"message"`
	Username        string `json:"username"`
	NamaShift       string `json:"nama_shift"`
	JamMasukShift   string `json:"jam_masuk_shift"`
	JamLogin        string `json:"jam_login"`
	StatusKehadiran string `json:"status_kehadiran"`
}

type Riwayat struct {
	IDPeminjaman  int    `json:"id_peminjaman"`
	Pustakawan    string `json:"pustakawan"`
	StatusPetugas string `json:"status_petugas"`
	TanggalPinjam string `json:"tanggal_pinjam"`
}
