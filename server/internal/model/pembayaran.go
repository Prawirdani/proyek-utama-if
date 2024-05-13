package model

type CreateMetodePembayaranRequest struct {
	TipePembayaran string `json:"tipePembayaran" validate:"required"`
	Metode         string `json:"metode" validate:"required"`
	Deskripsi      string `json:"deskripsi" validate:"required"`
}

type UpdateMetodePembayaranRequest struct {
	ID             int
	TipePembayaran string `json:"tipePembayaran" validate:"required"`
	Metode         string `json:"metode" validate:"required"`
	Deskripsi      string `json:"deskripsi" validate:"required"`
}

type PembayaranRequest struct {
	MetodePembayaranId int `json:"metodePembayaranId" validate:"required"`
	PesananId          int `json:"pesananId" validate:"required"`
}
