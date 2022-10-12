package messages

import (
	"gin-Cratos/request"
)

func newCargoMessage(req request.CargoRequest) {
	FormateaCampo3(req.Tipo)
	FormateaCampo4(req.Monto)
	FormateaCampo7()
	FormateaCampo11()
	FormateaCampo12()
	FormateaCampo13()
	FormateaCampo17()
	FormateaCampo18()
	FormateaCampo22(req.Tipo)
	FormateaCampo25()
	FormateaCampo32("12")
	FormateaCampo35(req.Tarjeta.Pan, req.Tarjeta.ExpAnio, req.Tarjeta.ExpMes)
	FormateaCampo37()
	FormateaCampo48(req.Afiliacion)
	FormateaCampo49(req.Moneda)
}

func SaleMessageContruct(sMti string, req request.CargoRequest) (request.SaleRequest, error) {

	newCargoMessage(req)
	res := request.SaleRequest{
		MTI:                 sMti,
		ProccesingCode:      Campo3,
		Amount:              Campo4,
		TransmitionDataTime: Campo7,
		SystemTrace:         Campo11,
		LocalTime:           Campo12,
		LocalDate:           Campo13,
		DateCapture:         Campo17,
		MerchantType:        Campo18,
		ServiceEntry:        Campo22,
		ServiceCondition:    Campo25,
		AcquiringIDCode:     Campo32,
		RetrievalReference:  Campo37,
		CardAcceptTerminal:  Campo41,
		RetailerData:        Campo48,
		CurrencyCode:        Campo49,
		TerminalData:        Campo60,
	}

	return res, nil
}
