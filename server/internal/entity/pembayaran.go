package entity

import "time"

type Pembayaran struct {
	Id              int              `json:"id"`
	PesananId       int              `json:"pesananId,omitempty"`
	Metode          MetodePembayaran `json:"metode"`
	Jumlah          int              `json:"jumlah"`
	WaktuPembayaran time.Time        `json:"waktuPembayaran"`
}

func NewPembayaran(pesanan Pesanan, metode MetodePembayaran) *Pembayaran {
	return &Pembayaran{
		PesananId:       pesanan.ID,
		Metode:          metode,
		Jumlah:          pesanan.Total,
		WaktuPembayaran: time.Now(),
	}
}
