package model

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
