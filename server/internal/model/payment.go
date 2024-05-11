package model

type CreateMetodePembayaranRequest struct {
	TipePembayaran string `json:"tipePembayaran"`
	Metode         string `json:"metode"`
	Deskripsi      string `json:"deskripsi"`
}

type UpdateMetodePembayaranRequest struct {
	ID             int
	TipePembayaran string `json:"tipePembayaran"`
	Metode         string `json:"metode"`
	Deskripsi      string `json:"deskripsi"`
}
