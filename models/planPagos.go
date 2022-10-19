package models

type PlanPagos struct {
	Tipo          string `json:"tipo" validate:"required"`
	Puntos        int    `json:"puntos" validate:"min=0"`
	Parcialidades int    `json:"parcialidades" validate:"min=0"`
	Diferido      int    `json:"diferido" validate:"min=0"`
}
