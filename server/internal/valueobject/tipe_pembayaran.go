package valueobject

import (
	"strings"

	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

type TipePembayaran string

const (
	TipePembayaranCash   TipePembayaran = "CASH"
	TipePembayaranCard   TipePembayaran = "CARD"
	TipePembayaranMobile TipePembayaran = "MOBILE"
)

var (
	ErrorTipePembayaranInvalid = httputil.ErrBadRequest("Tipe pembayaran tidak valid!")
)

// Check if the tipe pembayaran is validm if valid assign the value to tp
func (tp *TipePembayaran) Valid(str string) error {
	v := strings.ToUpper(str)
	switch v {
	case string(TipePembayaranCash), string(TipePembayaranCard), string(TipePembayaranMobile):
		tp = (*TipePembayaran)(&v)
		return nil
	}
	return ErrorTipePembayaranInvalid
}
