package entity

import (
	"testing"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/internal/valueobject"
	"github.com/stretchr/testify/assert"
)

var meja = Meja{
	ID:     1,
	Nomor:  "M-1",
	Status: valueobject.StatusMejaTersedia,
}

var menus = []Menu{
	{
		ID:    1,
		Nama:  "Nasi Goreng",
		Harga: 15000,
	},
	{
		ID:    2,
		Nama:  "Mie Goreng",
		Harga: 12000,
	},
}

var note = "Pesanan note"
var detailRequest = []model.PesananMenuRequest{
	{MenuID: menus[0].ID, Kuantitas: 2},
	{MenuID: menus[1].ID, Kuantitas: 1},
}
var requestDineIn = model.PesananDineInRequest{
	NamaPelanggan: "Prawirdani",
	KasirID:       1,
	MejaID:        1,
	Catatan:       &note,
}

var requestTakeAway = model.PesananTakeAwayRequest{
	NamaPelanggan: "Prawirdani",
	KasirID:       1,
	Catatan:       &note,
}

func TestNewPesananDineIn(t *testing.T) {
	t.Run("NewPesananDineIn", func(t *testing.T) {
		pesanan := NewPesananDineIn(requestDineIn)

		assert.Equal(t, requestDineIn.NamaPelanggan, pesanan.NamaPelanggan)
		assert.Equal(t, requestDineIn.KasirID, pesanan.Kasir.ID)
		assert.Equal(t, requestDineIn.MejaID, pesanan.Meja.ID)
		assert.Equal(t, valueobject.TipePesananDineIn, pesanan.TipePesanan)
		assert.Equal(t, valueobject.StatusPesananDiProses, pesanan.StatusPesanan)
		assert.NotEmpty(t, pesanan.WaktuPesanan)
	})

	t.Run("NewPesananDineIn-AddDetail", func(t *testing.T) {
		pesanan := NewPesananDineIn(requestDineIn)

		expectedTotal := 0
		for i := 0; i < len(detailRequest); i++ {
			menu := menus[i]
			detail, err := NewDetailPesanan(menu, detailRequest[i].Kuantitas)
			expectedTotal += detail.Subtotal
			assert.Nil(t, err)
			pesanan.AddDetail(detail)

		}

		assert.NotZero(t, pesanan.Total)
		assert.Equal(t, expectedTotal, pesanan.Total)
	})
}

func TestNewPesananTakeAway(t *testing.T) {
	t.Run("NewPesananTakeAway", func(t *testing.T) {
		pesanan := NewPesananTakeAway(requestTakeAway)

		assert.Equal(t, requestTakeAway.NamaPelanggan, pesanan.NamaPelanggan)
		assert.Equal(t, requestTakeAway.KasirID, pesanan.Kasir.ID)
		assert.Nil(t, pesanan.Meja)
		assert.Equal(t, valueobject.TipePesananTakeAway, pesanan.TipePesanan)
		assert.Equal(t, valueobject.StatusPesananDiProses, pesanan.StatusPesanan)
		assert.NotEmpty(t, pesanan.WaktuPesanan)
	})

	t.Run("NewPesananTakeAway-AddDetail", func(t *testing.T) {
		pesanan := NewPesananTakeAway(requestTakeAway)

		expectedTotal := 0
		for i := 0; i < len(detailRequest); i++ {
			menu := menus[i]
			detail, err := NewDetailPesanan(menu, detailRequest[i].Kuantitas)
			expectedTotal += detail.Subtotal
			assert.Nil(t, err)
			pesanan.AddDetail(detail)
		}

		assert.NotZero(t, pesanan.Total)
		assert.Equal(t, expectedTotal, pesanan.Total)
	})
}
