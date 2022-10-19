package request

type CancelacionRequest struct {
	UUID  string `json:"uuid"`
	Monto string `json:"monto" validate:"required"`
}
