package model

type CreateMejaRequest struct {
	Nomor string `json:"nomor" validate:"required"`
}

type UpdateMejaRequest struct {
	ID    int    `json:"id" validate:"required"`
	Nomor string `json:"nomor" validate:"required"`
}
