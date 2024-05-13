package model

import "time"

type PesananMenuRequest struct {
	MenuID    int `json:"menuID" validate:"required"`
	Kuantitas int `json:"kuantitas" validate:"required,min=1"`
}

type PesananDineInRequest struct {
	NamaPelanggan string               `json:"namaPelanggan" validate:"required"`
	MejaID        int                  `json:"mejaID" validate:"required"`
	Menu          []PesananMenuRequest `json:"menu" validate:"required"`
	Catatan       *string              `json:"catatan"`
	KasirID       int
}

type PesananTakeAwayRequest struct {
	NamaPelanggan string               `json:"namaPelanggan" validate:"required"`
	Menu          []PesananMenuRequest `json:"menu" validate:"required"`
	Catatan       *string              `json:"catatan"`
	KasirID       int
}

type PesananResponse struct {
	ID            int                     `json:"id"`
	NamaPelanggan string                  `json:"namaPelanggan"`
	Kasir         string                  `json:"kasir"`
	Meja          string                  `json:"meja"`
	Tipe          string                  `json:"tipe"`
	Status        string                  `json:"status"`
	Catatan       *string                 `json:"catatan"`
	Detail        []PesananDetailResponse `json:"detail"`
	Total         int                     `json:"total"`
	WaktuPesanan  time.Time               `json:"waktuPesanan"`
}

type PesananDetailResponse struct {
	ID        int    `json:"id"`
	NamaMenu  string `json:"namaMenu"`
	HargaMenu int    `json:"hargaMenu"`
	Kuantitas int    `json:"kuantitas"`
	Subtotal  int    `json:"subtotal"`
}
