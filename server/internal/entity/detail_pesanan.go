package entity

type DetailPesanan struct {
	ID        int  `json:"id"`
	PesananID int  `json:"pesananId,omitempty"`
	Menu      Menu `json:"menu"`
	Kuantitas int  `json:"kuantitas"`
	Subtotal  int  `json:"subtotal"`
}

func (d *DetailPesanan) CalculateSubtotal() {
	d.Subtotal = d.Menu.Harga * d.Kuantitas
}
