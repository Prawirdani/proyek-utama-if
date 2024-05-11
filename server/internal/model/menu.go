package model

type CreateKategoriMenuRequest struct {
	Nama string `json:"nama" validate:"required"`
}

type UpdateKategoriMenuRequest struct {
	ID   int    `validate:"required"`
	Nama string `json:"nama" validate:"required"`
}

type CreateMenuRequest struct {
	Nama       string  `json:"nama" validate:"required"`
	Deskripsi  *string `json:"deskripsi"`
	Harga      int     `json:"harga" validate:"required,min=1"`
	KategoriId int     `json:"kategoriId" validate:"required"`
	ImageName  *string
}

type UpdateMenuRequest struct {
	ID         int     `validate:"required"`
	Nama       string  `json:"nama" validate:"required"`
	Deskripsi  *string `json:"deskripsi"`
	Harga      int     `json:"harga" validate:"required,min=1"`
	KategoriId int     `json:"kategoriId" validate:"required"`
	ImageName  *string
}
