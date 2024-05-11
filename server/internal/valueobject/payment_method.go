package valueobject

import (
	"strings"
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorInvalidTipePembayaran = httputil.ErrBadRequest("Tipe pembayaran tidak valid!")
)

type TipePembayaran string

const (
	TipePembayaranCash   = "CASH"
	TipePembayaranCard   = "CARD"
	TipePembayaranMobile = "MOBILE"
)

type MetodePembayaran struct {
	ID             int            `json:"id"`
	TipePembayaran TipePembayaran `json:"tipePembayaran"`
	Metode         string         `json:"metode"`
	Deskripsi      string         `json:"deskripsi"`
	DeletedAt      *time.Time     `json:"-"`
}

func (m *MetodePembayaran) Assign(request model.UpdateMetodePembayaranRequest) {
	m.TipePembayaran = TipePembayaran(request.TipePembayaran)
	m.Metode = request.Metode
	m.Deskripsi = request.Deskripsi
}

func NewMetodePembayaran(request model.CreateMetodePembayaranRequest) (MetodePembayaran, error) {
	if !isValidTipePembayaran(request.TipePembayaran) {
		return MetodePembayaran{}, ErrorInvalidTipePembayaran
	}

	return MetodePembayaran{
		TipePembayaran: TipePembayaran(request.TipePembayaran),
		Metode:         request.Metode,
		Deskripsi:      request.Deskripsi,
	}, nil
}

func isValidTipePembayaran(str string) bool {
	str = strings.ToUpper(str)
	switch str {
	case string(TipePembayaranCash), string(TipePembayaranCard), string(TipePembayaranMobile):
		return true
	}
	return false
}
