package entity

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/prawirdani/golang-restapi/internal/model"
	"github.com/prawirdani/golang-restapi/pkg/httputil"
)

var (
	ErrorKategoriDuplicate = httputil.ErrConflict("Kategori dengan nama tersebut sudah ada!")
	ErrorKategoriNotFound  = httputil.ErrNotFound("Kategori tidak ditemukan.")
	ErrorMenuNotFound      = httputil.ErrNotFound("Menu tidak ditemukan.")
)

func capitalizeFirstLetter(str string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(str[0:1]), str[1:])
}

type KategoriMenu struct {
	ID        int        `json:"id,omitempty"`
	Nama      string     `json:"nama,omitempty"`
	DeletedAt *time.Time `json:"-"`
}

func (k *KategoriMenu) SetDeletedAt() {
	t := time.Now()
	k.DeletedAt = &t
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
	ID        int           `json:"id,omitempty"`
	Nama      string        `json:"nama"`
	Deskripsi string        `json:"deskripsi,omitempty"`
	Harga     int           `json:"harga"`
	Kategori  *KategoriMenu `json:"kategori,omitempty"`
	Url       *string       `json:"url,omitempty"`
	DeletedAt *time.Time    `json:"-"`
	CreatedAt *time.Time    `json:"createdAt,omitempty"`
	UpdatedAt *time.Time    `json:"updatedAt,omitempty"`
}

func (m *Menu) ScanRow(r Row) error {
	kategoriID := sql.NullInt64{}
	kategoriNama := sql.NullString{}
	kategoriDeletedAt := sql.NullTime{}

	err := r.Scan(&m.ID,
		&m.Nama,
		&m.Deskripsi,
		&m.Harga,
		&m.Url,
		&m.DeletedAt,
		&m.CreatedAt,
		&m.UpdatedAt,
		&kategoriID,
		&kategoriNama,
		&kategoriDeletedAt,
	)
	if err != nil {
		return err
	}

	if kategoriID.Valid {
		m.Kategori = &KategoriMenu{
			ID:        int(kategoriID.Int64),
			Nama:      kategoriNama.String,
			DeletedAt: &kategoriDeletedAt.Time,
		}
	}
	return nil
}

func (m *Menu) Assign(request model.UpdateMenuRequest) {
	m.Nama = capitalizeFirstLetter(request.Nama)
	m.Deskripsi = *request.Deskripsi
	m.Harga = request.Harga
	m.Kategori.ID = request.KategoriId
	if request.ImageName != nil {
		m.Url = request.ImageName
	}
}

func (m *Menu) SetDeletedAt() {
	t := time.Now()
	m.DeletedAt = &t
}

func (m *Menu) DeleteImage() error {
	if m.Url != nil {
		filename := *m.Url
		m.Url = nil
		return httputil.DeleteUpload(filename)
	}
	return nil
}

func NewMenu(request model.CreateMenuRequest) Menu {
	return Menu{
		Nama:      capitalizeFirstLetter(request.Nama),
		Deskripsi: *request.Deskripsi,
		Harga:     request.Harga,
		Kategori: &KategoriMenu{
			ID: request.KategoriId,
		},
		Url: request.ImageName,
	}
}
