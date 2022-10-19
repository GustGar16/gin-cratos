package messages

import (
	"gin-Cratos/classes"
	"gin-Cratos/request"
	"gin-Cratos/services/database"
)

func newCaptureMessage(req request.CapturaRequest) {
	//FormateaCampo3(req.Tipo)
	FormateaCampo4(req.Monto)
	FormateaCampo7()
	FormateaCampo11()
	FormateaCampo12()
	//FormateaCampo13()
	//FormateaCampo17()
	//FormateaCampo18()
	//FormateaCampo22(req.Tipo)
	//FormateaCampo25()
	//FormateaCampo32("12")
	//FormateaCampo35(req.Tarjeta.Pan, req.Tarjeta.ExpAnio, req.Tarjeta.ExpMes)
	//FormateaCampo37()
	//FormateaCampo48(req.Afiliacion)
	//FormateaCampo42()
	//FormateaCampo43()
	//FormateaCampo49(req.Moneda)
	//FormateaCampo54(req.Monto)
	//FormateaCampo61()
	//FormateaCampo63()
	//FormateaCampo90()
	//FormateaCampo95()
	//FormateaCampo100()
	//FormateaCampo120()
	//FormateaCampo121()
	//FormateaCampo123()
	//FormateaCampo125()
	//FormateaCampo126()
}

func CaptureMessageContruct(sMti string, req request.CapturaRequest) (classes.CaptureMessage, error) {
	var captureData classes.CaptureMessage
	respuesta, err := database.GetMessageRequest(req.UUID)
	if err != nil {
		return captureData, err
	}
	newCaptureMessage(req)
	res := classes.CaptureMessage{
		MTI:                 sMti,
		ProccesingCode:      Campo3,
		Amount:              Campo4,
		TransmitionDataTime: Campo7,
		SystemTrace:         Campo11,
		LocalTime:           Campo12,
		LocalDate:           Campo13,
		DateCapture:         Campo17,
		MerchantType:        respuesta.MerchantType,
		ServiceEntry:        Campo22,
		ServiceCondition:    Campo25,
		AcquiringIDCode:     Campo32,
		RetrievalReference:  Campo37,
		CardAcceptTerminal:  Campo41,
		//CardAcceptorCode: Campo42,
		//CardAcceptorName: Campo43,
		RetailerData: Campo48,
		CurrencyCode: Campo49,
		//AdditionalAmounts: Campo54,
		TerminalData: Campo60,
		//CardIssuer: Campo61,
		//PostalCode: Campo62,
		//AdditionalData: Campo63,
		//OriginalData: Campo90,
		//ReplacementAmounts: Campo95,
		//ReceivingIDCode: Campo100,
		//TerminalAddress: Campo120,
		//AuthIndicators: Campo121,
		//PosInvoiceData: Campo123,
		//PosSettlementData: Campo125,
		//PosPreauthData: Campo126,
	}

	return res, nil
}
