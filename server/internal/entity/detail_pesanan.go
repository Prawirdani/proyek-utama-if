package entity

import "github.com/prawirdani/golang-restapi/pkg/httputil"

type DetailPesanan struct {
	ID        int  `json:"id"`
	PesananID int  `json:"pesananId,omitempty"`
	Menu      Menu `json:"menu,omitempty"`
	Kuantitas int  `json:"kuantitas"`
	Subtotal  int  `json:"subtotal"`
}

var (
	ErrorHargaMenuTidakValid = httputil.ErrBadRequest("Harga menu tidak valid!")
	ErrorKuantitasMenuMin    = httputil.ErrBadRequest("Kuantitas menu minimal 1!")
)

func NewDetailPesanan(menu Menu, kuantitas int) (DetailPesanan, error) {
	if menu.Harga <= 0 {
		return DetailPesanan{}, ErrorHargaMenuTidakValid
	}

	if kuantitas < 1 {
		return DetailPesanan{}, ErrorKuantitasMenuMin
	}

	return DetailPesanan{
		Menu:      menu,
		Kuantitas: kuantitas,
		Subtotal:  menu.Harga * kuantitas,
	}, nil
}
