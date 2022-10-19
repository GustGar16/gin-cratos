package request

import "gin-Cratos/models"

type CargoRequest struct {
	UUID       string           `json:"uuid" validate:"required,len=36"`
	Monto      string           `json:"monto" validate:"required"`
	Tipo       string           `json:"tipo" validate:"required"`
	Moneda     string           `json:"moneda" validate:"required"`
	Afiliacion string           `json:"afiliacion" validate:"required"`
	Tarjeta    models.Tarjeta   `json:"tarjeta"`
	Plan       models.PlanPagos `json:"plan"`
}
