package entity

import "github.com/prawirdani/golang-restapi/internal/model"

type Receipt struct {
	Pesanan    model.PesananResponse `json:"pesanan"`
	Pembayaran Pembayaran            `json:"pembayaran"`
}

func NewReceipt(pesanan Pesanan, pembayaran Pembayaran) *Receipt {
	return &Receipt{
		Pesanan:    pesanan.ToResponse(),
		Pembayaran: pembayaran,
	}
}
