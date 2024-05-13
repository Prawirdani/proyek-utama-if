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

func NewPesananDineIn(req model.PesananDineInRequest, meja *Meja) (Pesanan, error) {
	if err := meja.SetTerisi(); err != nil {
		return Pesanan{}, err
	}

	return Pesanan{
		NamaPelanggan: req.NamaPelanggan,
		Kasir: User{
			ID: req.KasirID,
		},
		Meja:          meja,
		TipePesanan:   valueobject.TipePesananDineIn,
		StatusPesanan: valueobject.StatusPesananDiProses,
		Catatan:       req.Catatan,
		WaktuPesanan:  time.Now(),
	}, nil
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

func (p Pesanan) IsDineIn() bool {
	return p.TipePesanan == valueobject.TipePesananDineIn
}

func (p Pesanan) IsBatal() bool {
	return p.StatusPesanan == valueobject.StatusPesananBatal
}

func (p Pesanan) IsSelesai() bool {
	return p.StatusPesanan == valueobject.StatusPesananSelesai
}

func (p Pesanan) IsDisajikan() bool {
	return p.StatusPesanan == valueobject.StatusPesananDisajikan
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

// Set pesanan to selesai, and set meja to tersedia (if dine in)
func (p *Pesanan) Selesaikan() error {
	if p.IsSelesai() {
		return ErrorPesananAlreadySelesai
	}

	if p.IsDineIn() {
		if err := p.Meja.SetTersedia(); err != nil {
			return err
		}
	}

	p.StatusPesanan = valueobject.StatusPesananSelesai
	return nil
}

// Set Pesanan status to batal, and set meja to tersedia (if dine in)
func (p *Pesanan) Batalkan() error {
	if p.IsBatal() {
		return ErrorPesananAlreadyBatal
	}
	if p.IsDineIn() {
		if err := p.Meja.SetTersedia(); err != nil {
			return err
		}
	}
	p.StatusPesanan = valueobject.StatusPesananBatal
	return nil
}

func (p *Pesanan) Sajikan() error {
	if p.IsDisajikan() {
		return ErrorPesananAlreadyDisajikan
	}
	p.StatusPesanan = valueobject.StatusPesananDisajikan
	return nil
}

func (p *Pesanan) ScanRow(r Row) error {
	mejaId := sql.NullInt64{}
	mejaNomor := sql.NullString{}
	mejaStatus := sql.NullString{}
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
		&mejaStatus,
	)
	if err != nil {
		return err
	}
	if mejaId.Valid {
		p.Meja = &Meja{
			ID:     int(mejaId.Int64),
			Nomor:  mejaNomor.String,
			Status: valueobject.StatusMeja(mejaStatus.String),
		}
	}

	return nil
}
