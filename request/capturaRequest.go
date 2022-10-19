package request

import "gin-Cratos/models"

type CapturaRequest struct {
	UUID    string         `json:"uuid"`
	Monto   string         `json:"monto" validate:"required"`
	Tarjeta models.Tarjeta `json:"tarjeta"`
}
