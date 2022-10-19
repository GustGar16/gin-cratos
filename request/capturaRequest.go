package request

type CapturaRequest struct {
	UUID  string `json:"uuid"`
	Monto string `json:"monto" validate:"required"`
}
