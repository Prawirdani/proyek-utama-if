package entity

import (
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
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
	Catatan       *string                   `json:"catatan,omitempty"`
	Total         int                       `json:"total"`
	WaktuPesanan  time.Time                 `json:"waktuPesanan"`
}

func NewPesananDineIn(req model.PesananDineInRequest) Pesanan {
	return Pesanan{
		NamaPelanggan: req.NamaPelanggan,
		Kasir: User{
			ID: req.KasirID,
		},
		Meja: &Meja{
			ID: req.MejaID,
		},
		TipePesanan:   valueobject.TipePesananDineIn,
		StatusPesanan: valueobject.StatusPesananDiProses,
		Catatan:       req.Catatan,
		WaktuPesanan:  time.Now(),
	}
}

func NewPesananTakeAway(req model.PesananTakeAwayRequest) Pesanan {
	return Pesanan{
		NamaPelanggan: req.NamaPelanggan,
		Kasir: User{
			ID: req.KasirID,
		},
		TipePesanan:   valueobject.TipePesananTakeAway,
		StatusPesanan: valueobject.StatusPesananDiProses,
		Catatan:       req.Catatan,
		WaktuPesanan:  time.Now(),
	}
}

func (p *Pesanan) AddDetail(detail DetailPesanan) {
	p.Detail = append(p.Detail, detail)
	p.Total += detail.Subtotal
}

func (p *Pesanan) Selesaikan() {
	p.StatusPesanan = valueobject.StatusPesananSelesai
}

func (p *Pesanan) Batalkan() {
	p.StatusPesanan = valueobject.StatusPesananBatal
}

func (p *Pesanan) Sajikan() {
	p.StatusPesanan = valueobject.StatusPesananDisajikan
}
