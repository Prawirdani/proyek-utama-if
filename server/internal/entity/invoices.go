package entity

import "github.com/prawirdani/golang-restapi/internal/model"

type Invoice struct {
	Pesanan    model.PesananResponse `json:"pesanan"`
	Pembayaran Pembayaran            `json:"pembayaran"`
}

func NewInvoice(pesanan Pesanan, pembayaran Pembayaran) *Invoice {
	return &Invoice{
		Pesanan:    pesanan.ToResponse(),
		Pembayaran: pembayaran,
	}
}
