package request

type ReembolsoRequest struct {
	UUID  string `json:"uuid"`
	Monto string `json:"monto" validate:"required"`
}
