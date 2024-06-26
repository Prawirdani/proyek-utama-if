package entity

import (
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/valueobject"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorMejaDuplicate       = httputil.ErrConflict("Meja dengan nomor tersebut sudah ada!")
	ErrorMejaNotFound        = httputil.ErrConflict("Meja tidak ditemukan!")
	ErrorMejaTidakTersedia   = httputil.ErrConflict("Meja tidak tersedia!")
	ErrorMejaAlreadyTerisi   = httputil.ErrConflict("Meja sudah terisi!")
	ErrorMejaAlreadyReserved = httputil.ErrConflict("Meja sudah di-reserved!")
	ErrorMejaAlreadyTersedia = httputil.ErrConflict("Meja sudah tersedia!")
)

type Meja struct {
	ID        int                    `json:"id"`
	Nomor     string                 `json:"nomor"`
	Status    valueobject.StatusMeja `json:"status,omitempty"`
	DeletedAt *time.Time             `json:"-"`
}

func (m Meja) Tersedia() bool {
	return m.Status == valueobject.StatusMejaTersedia
}

func (m Meja) Reserved() bool {
	return m.Status == valueobject.StatusMejaReserved
}

func (m Meja) Terisi() bool {
	return m.Status == valueobject.StatusMejaTerisi
}

// Set Meja status to Terisi and also check if Meja is Tersedia
func (m *Meja) SetTerisi() error {
	if !m.Tersedia() {
		return ErrorMejaTidakTersedia
	}

	if m.Terisi() {
		return ErrorMejaAlreadyTerisi
	}

	m.Status = valueobject.StatusMejaTerisi
	return nil
}

func (m *Meja) SetTersedia() error {
	if m.Tersedia() {
		return ErrorMejaAlreadyTersedia
	}
	m.Status = valueobject.StatusMejaTersedia
	return nil
}

func (m *Meja) SetReserved() error {
	if !m.Tersedia() {
		return ErrorMejaTidakTersedia
	}
	if m.Reserved() {
		return ErrorMejaAlreadyReserved
	}
	m.Status = valueobject.StatusMejaReserved
	return nil
}

func (m *Meja) ScanRow(row Row) error {
	return row.Scan(&m.ID, &m.Nomor, &m.Status, &m.DeletedAt)
}

func (m *Meja) Assign(request model.UpdateMejaRequest) {
	m.Nomor = request.Nomor
}

func (m *Meja) SetDeletedAt() {
	now := time.Now()
	m.DeletedAt = &now
}

func NewMeja(request model.CreateMejaRequest) *Meja {
	return &Meja{
		Nomor:  request.Nomor,
		Status: valueobject.StatusMejaTersedia,
	}
}
