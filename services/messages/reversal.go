package messages

import (
	"gin-Cratos/request"
	"gin-Cratos/services/database"
)

func newCancelacionMessage(req request.SaleRequest) {
	FormateaCampo7()
	FormateaCampo11()
	FormateaCampo12()
	FormateaCampo13()
	FormateaCampo17()
	FormateaCampo18()
	FormateaCampo32("12")
}

func ReversalMessageContruct(sMti string, transaccionId string) (request.ReversalRequest, error) {
	var reversalData request.ReversalRequest
	respuesta, err := database.GetMessageRequest(transaccionId)
	if err != nil {
		return reversalData, err
	}
	newCancelacionMessage(respuesta)

	reversalData.ProccesingCode = respuesta.ProccesingCode
	reversalData.Amount = respuesta.Amount
	reversalData.TransmitionDataTime = Campo7
	reversalData.SystemTrace = Campo11
	reversalData.LocalTime = Campo12
	reversalData.LocalDate = Campo13
	reversalData.DateCapture = Campo17
	reversalData.MerchantType = Campo18
	reversalData.ServiceEntry = respuesta.ServiceEntry
	reversalData.ServiceCondition = respuesta.ServiceCondition
	reversalData.AcquiringIDCode = Campo32

	return reversalData, nil

}
