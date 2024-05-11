package valueobject

import (
	"fmt"
	"strings"
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorTipePembayaranInvalid     = httputil.ErrBadRequest("Tipe pembayaran tidak valid!")
	ErrorMetodePembayaranDuplicate = httputil.ErrBadRequest("Metode pembayaran yang sama sudah ada!")
	ErrorMetodePembayaranNotFound  = httputil.ErrNotFound("Metode pembayaran tidak ditemukan!")
)

type tipePembayaran string

const (
	TipePembayaranCash   = "CASH"
	TipePembayaranCard   = "CARD"
	TipePembayaranMobile = "MOBILE"
)

type MetodePembayaran struct {
	ID             int            `json:"id"`
	TipePembayaran tipePembayaran `json:"tipePembayaran"`
	Metode         string         `json:"metode"`
	Deskripsi      string         `json:"deskripsi"`
	DeletedAt      *time.Time     `json:"-"`
}

func (m *MetodePembayaran) ScanRow(row Row) error {
	return row.Scan(&m.ID, &m.TipePembayaran, &m.Metode, &m.Deskripsi, &m.DeletedAt)
}

func (m *MetodePembayaran) Assign(request model.UpdateMetodePembayaranRequest) error {
	tp := strings.ToUpper(request.TipePembayaran)
	if !isValidTipePembayaran(tp) {
		return ErrorTipePembayaranInvalid
	}
	m.TipePembayaran = tipePembayaran(tp)
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

func NewMetodePembayaran(request model.CreateMetodePembayaranRequest) (MetodePembayaran, error) {
	tp := strings.ToUpper(request.TipePembayaran)
	if !isValidTipePembayaran(tp) {
		return MetodePembayaran{}, ErrorTipePembayaranInvalid
	}

	return MetodePembayaran{
		TipePembayaran: tipePembayaran(tp),
		Metode:         request.Metode,
		Deskripsi:      request.Deskripsi,
	}, nil
}

func isValidTipePembayaran(str string) bool {
	switch str {
	case string(TipePembayaranCash), string(TipePembayaranCard), string(TipePembayaranMobile):
		return true
	}
	return false
}
