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
	FormateaCampo13()
	FormateaCampo17()
	FormateaCampo25()
	FormateaCampo32("12")
	FormateaCampo35(req.Tarjeta.Pan, req.Tarjeta.ExpAnio, req.Tarjeta.ExpMes)
	FormateaCampo37()
	//FormateaCampo42()
	//FormateaCampo43()
	//FormateaCampo54(req.Monto)
	//FormateaCampo61()
	//FormateaCampo62()
	//FormateaCampo63()
	//FormateaCampo90()
	//FormateaCampo95()
	//FormateaCampo100()
	//FormateaCampo102()
	//FormateaCampo120()
	//FormateaCampo121()
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
		ServiceEntry:        respuesta.ServiceEntry,
		ServiceCondition:    Campo25,
		AcquiringIDCode:     Campo32,
		Track2Data:          Campo35,
		RetrievalReference:  Campo37,
		AuthIDResponse:      "", //TODO
		ResponseCode:        "", //TODO
		CardAcceptTerminal:  Campo41,
		//CardAcceptorCode: Campo42,
		//CardAcceptorName: Campo43,
		AdditionalResponse: "",
		Track1Data:         "",
		RetailerData:       respuesta.RetailerData,
		CurrencyCode:       respuesta.CurrencyCode,
		AdditionalAmounts:  respuesta.AdditionalAmounts,
		TerminalData:       Campo60,
		//CardIssuer: Campo61,
		PostalCode: Campo62,
		//AdditionalData: Campo63,
		//OriginalData: Campo90,
		//ReplacementAmounts: Campo95,
		//ReceivingIDCode: Campo100,
		AccountID1: "",
		//TerminalAddress: Campo120,
		//AuthIndicators: Campo121,
		CardIssuerID:      "",
		PosInvoiceData:    respuesta.PosInvoiceData,
		PosSettlementData: respuesta.PosSettlementData,
		PosPreauthData:    respuesta.PosPreauthData,
	}

	return res, nil
}
