package entity

import (
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorDuplicateNomorMeja = httputil.ErrConflict("Meja dengan nomor tersebut sudah ada!")
	ErrorMejaNotFound       = httputil.ErrConflict("Meja tidak ditemukan!")
)

type StatusMeja string

const (
	statusMejaTersedia = "Tersedia"
	statusMejaTerisi   = "Terisi"
	statusMejaReserved = "Reserved"
)

type Meja struct {
	ID        int        `json:"id"`
	Nomor     string     `json:"nomor"`
	Status    StatusMeja `json:"status"`
	DeletedAt *time.Time `json:"-"`
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

func NewMeja(request model.CreateMejaRequest) Meja {
	return Meja{
		Nomor:  request.Nomor,
		Status: statusMejaTersedia,
	}
}
