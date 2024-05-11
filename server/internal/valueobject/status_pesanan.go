package valueobject

type StatusPesanan string

const (
	StatusPesananDiProses  StatusPesanan = "Diproses"
	StatusPesananDisajikan StatusPesanan = "Disajikan"
	StatusPesananBatal     StatusPesanan = "Batal"
	StatusPesananSelesai   StatusPesanan = "Selesai"
)
