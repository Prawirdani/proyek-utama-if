package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorKategoriExists   = httputil.ErrConflict("Kategori already exists.")
	ErrorKategoriNotFound = httputil.ErrNotFound("Kategori not found.")
	ErrorMenuNotFound     = httputil.ErrNotFound("Menu not found.")
)

func capitalizeFirstLetter(str string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(str[0:1]), str[1:])
}

type KategoriMenu struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func (k *KategoriMenu) SetDeletedAt() {
	t := time.Now()
	k.DeletedAt = &t
}

func (k KategoriMenu) NameEqual(nama string) bool {
	return strings.EqualFold(k.Nama, nama)
}

func (k *KategoriMenu) ScanRow(r Row) error {
	return r.Scan(&k.ID, &k.Nama, &k.DeletedAt)
}

func (k *KategoriMenu) Assign(request model.UpdateKategoriMenuRequest) {
	k.Nama = capitalizeFirstLetter(request.Nama)
}

func NewKategoriMenu(request model.CreateKategoriMenuRequest) KategoriMenu {
	return KategoriMenu{
		Nama: capitalizeFirstLetter(request.Nama),
	}
}

type Menu struct {
	ID        int          `json:"id"`
	Nama      string       `json:"nama"`
	Deskripsi string       `json:"deskripsi"`
	Harga     int          `json:"harga"`
	Kategori  KategoriMenu `json:"kategori"`
	Url       *string      `json:"url"`
	DeletedAt *time.Time   `json:"deletedAt"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

func (m *Menu) ScanRow(r Row) error {
	return r.Scan(&m.ID,
		&m.Nama,
		&m.Deskripsi,
		&m.Harga,
		&m.Url,
		&m.DeletedAt,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.Kategori.ID,
		&m.Kategori.Nama,
		&m.Kategori.DeletedAt,
	)
}

func (m *Menu) Assign(request model.UpdateMenuRequest) {
	m.Nama = capitalizeFirstLetter(request.Nama)
	m.Deskripsi = *request.Deskripsi
	m.Harga = request.Harga
	m.Kategori.ID = request.KategoriId
}

func (m *Menu) SetDeletedAt() {
	t := time.Now()
	m.DeletedAt = &t
}

func NewMenu(request model.CreateMenuRequest) Menu {
	return Menu{
		Nama:      capitalizeFirstLetter(request.Nama),
		Deskripsi: *request.Deskripsi,
		Harga:     request.Harga,
		Kategori: KategoriMenu{
			ID: request.KategoriId,
		},
	}
}
