package entity

import "time"

type Pembayaran struct {
	Id              int              `json:"id"`
	PesananId       int              `json:"pesanan_id,omitempty"`
	Metode          MetodePembayaran `json:"metode"`
	Jumlah          int              `json:"jumlah"`
	WaktuPembayaran time.Time        `json:"waktu_pembayaran"`
}
