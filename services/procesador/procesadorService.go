package procesador

import (
	"gin-Cratos/classes"
	"gin-Cratos/request"
	"gin-Cratos/services/messages"
	"gin-Cratos/services/schema"
)

var mti string

func SetMTI(value string) bool {
	mti = value
	return true
}
func init() {
	schema.New()
	//Validar parametros de entrada
	//generar objeto request (si aplica)
	//Validar formato de cada parametro del request
	//Crear o consultar registro de transaccion en bd
	//definir mti del mensaje
	//Definir campos del mensaje
	//Definir el bitmap del mensaje
	//concatenar y codificar mensaje
	//Transmitir y recibir respuesta de mensaje
}

func Sale(request request.CargoRequest) (classes.SaleMessage, error) {
	SetMTI("0200")
	message, err := messages.SaleMessageContruct(mti, request)

	return message, err

}

func Reversal(request request.CancelacionRequest) (classes.ReversalMessage, error) {
	SetMTI("0420")
	message, err := messages.ReversalMessageContruct(mti, request)

	return message, err

}

func Void(request request.CancelacionRequest) (classes.ReversalMessage, error) {
	SetMTI("0420")
	message, err := messages.VoidMessageContruct(mti, request)

	return message, err

}

func Refund(request request.ReembolsoRequest) {
	//setMTI("0200")

}

func Capture(request request.CapturaRequest) (classes.CaptureMessage, error) {
	SetMTI("0220")
	message, err := messages.CaptureMessageContruct(mti, request)

	return message, err
}
