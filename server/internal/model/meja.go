package model

type CreateMejaRequest struct {
	Nomor string `json:"nomor" validate:"required"`
}

type UpdateMejaRequest struct {
	ID    int
	Nomor string `json:"nomor" validate:"required"`
}
