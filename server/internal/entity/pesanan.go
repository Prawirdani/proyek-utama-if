package entity

import (
	"database/sql"
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/valueobject"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorPesananNotFound         = httputil.ErrNotFound("Pesanan tidak ditemukan")
	ErrorPesananDetailNotExists  = httputil.ErrNotFound("Detail pesanan tidak ditemukan")
	ErrorPesananAlreadyBatal     = httputil.ErrBadRequest("Pesanan sudah dibatalkan")
	ErrorPesananAlreadySelesai   = httputil.ErrBadRequest("Pesanan sudah selesai")
	ErrorPesananAlreadyDisajikan = httputil.ErrBadRequest("Pesanan sudah disajikan")
)

type Pesanan struct {
	ID            int                       `json:"id"`
	NamaPelanggan string                    `json:"namaPelanggan"`
	Kasir         User                      `json:"kasir"`
	Meja          *Meja                     `json:"meja"`
	TipePesanan   valueobject.TipePesanan   `json:"tipe"`
	StatusPesanan valueobject.StatusPesanan `json:"status"`
	Detail        []DetailPesanan           `json:"detailPesanan"`
	Catatan       *string                   `json:"catatan"`
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

func (p *Pesanan) RemoveDetail(detailID int) error {
	for i, d := range p.Detail {
		if d.ID == detailID {
			p.Total -= d.Subtotal
			p.Detail = append(p.Detail[:i], p.Detail[i+1:]...)
			return nil
		}
	}
	return ErrorPesananDetailNotExists
}

func (p *Pesanan) Selesaikan() error {
	if p.StatusPesanan == valueobject.StatusPesananSelesai {
		return ErrorPesananAlreadySelesai
	}
	p.StatusPesanan = valueobject.StatusPesananSelesai
	return nil
}

func (p *Pesanan) Batalkan() error {
	if p.StatusPesanan == valueobject.StatusPesananBatal {
		return ErrorPesananAlreadyBatal
	}
	p.StatusPesanan = valueobject.StatusPesananBatal
	return nil
}

func (p *Pesanan) Sajikan() error {
	if p.StatusPesanan == valueobject.StatusPesananDisajikan {
		return ErrorPesananAlreadyDisajikan
	}
	p.StatusPesanan = valueobject.StatusPesananDisajikan
	return nil
}

func (p *Pesanan) ScanRow(r Row) error {
	mejaId := sql.NullInt64{}
	mejaNomor := sql.NullString{}
	err := r.Scan(
		&p.ID,
		&p.NamaPelanggan,
		&p.Total,
		&p.TipePesanan,
		&p.StatusPesanan,
		&p.Catatan,
		&p.WaktuPesanan,
		&p.Kasir.Nama,
		&mejaId,
		&mejaNomor,
	)
	if err != nil {
		return err
	}
	if mejaId.Valid {
		p.Meja = &Meja{
			ID:    int(mejaId.Int64),
			Nomor: mejaNomor.String,
		}
	}

	return nil
}
