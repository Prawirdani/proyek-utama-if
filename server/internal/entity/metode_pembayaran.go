package entity

import (
	"fmt"
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/valueobject"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorMetodePembayaranDuplicate = httputil.ErrBadRequest("Metode pembayaran yang sama sudah ada!")
	ErrorMetodePembayaranNotFound  = httputil.ErrNotFound("Metode pembayaran tidak ditemukan!")
)

type MetodePembayaran struct {
	ID             int                        `json:"id"`
	TipePembayaran valueobject.TipePembayaran `json:"tipePembayaran"`
	Metode         string                     `json:"metode"`
	Deskripsi      string                     `json:"deskripsi"`
	DeletedAt      *time.Time                 `json:"-"`
}

func (m *MetodePembayaran) ScanRow(row Row) error {
	return row.Scan(&m.ID, &m.TipePembayaran, &m.Metode, &m.Deskripsi, &m.DeletedAt)
}

func (m *MetodePembayaran) Assign(request model.UpdateMetodePembayaranRequest) error {
	if err := m.TipePembayaran.Valid(request.TipePembayaran); err != nil {
		return err
	}
	m.Metode = request.Metode
	m.Deskripsi = request.Deskripsi

	return nil
}

func (m *MetodePembayaran) SetDeletedAt() {
	now := time.Now()
	m.DeletedAt = &now
	// At time stamp, for bypassing unique constraint
	m.Metode = fmt.Sprintf("%s(%s)", m.Metode, now.Format("02-01-2006"))
}

func NewMetodePembayaran(request model.CreateMetodePembayaranRequest) (*MetodePembayaran, error) {
	mp := new(MetodePembayaran)
	if err := mp.TipePembayaran.Valid(request.TipePembayaran); err != nil {
		return nil, err
	}
	mp.Metode = request.Metode
	mp.Deskripsi = request.Deskripsi

	return mp, nil
}
