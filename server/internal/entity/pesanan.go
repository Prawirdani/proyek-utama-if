package entity

import (
	"database/sql"
	"time"

	"github.com/prawirdani/golang-restapi/internal/valueobject"
)

type Pesanan struct {
	ID            int                       `json:"id"`
	NamaPelanggan string                    `json:"namaPelanggan"`
	Kasir         User                      `json:"kasir"`
	Meja          *Meja                     `json:"meja,omitempty"`
	TipePesanan   valueobject.TipePesanan   `json:"tipe"`
	StatusPesanan valueobject.StatusPesanan `json:"status"`
	Detail        []DetailPesanan           `json:"detail"`
	Pembayaran    *Pembayaran               `json:"pembayaran"`
	Catatan       sql.NullString            `json:"catatan,omitempty"`
	WaktuPesanan  time.Time                 `json:"waktuPesanan"`
}

func (p *Pesanan) CalculateTotal() int {
	total := 0
	for _, detail := range p.Detail {
		total += detail.Subtotal
	}
	return total
}
