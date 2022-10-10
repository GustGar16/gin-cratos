package models

type Tarjeta struct {
	Pan     string `json:"pan" validate:"required,len=16"`
	Cvv     string `json:"cvv2" validate:"required,len=3"`
	ExpMes  string `json:"expiracion_mes" validate:"required,len=2"`
	ExpAnio string `json:"expiracion_anio" validate:"required,len=2"`
}
