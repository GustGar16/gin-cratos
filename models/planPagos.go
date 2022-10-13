package models

type PlanPagos struct {
	TIPO          string `json:"tipo" validate:"required"`
	PUNTOS        int    `json:"puntos" validate:"min=0"`
	PARCIALIDADES int    `json:"parcialidades" validate:"min=0"`
	DIFERIDO      int    `json:"diferido" validate:"min=0"`
}
