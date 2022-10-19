package messages

import (
	"gin-Cratos/classes"
	"gin-Cratos/request"
	"gin-Cratos/services/database"
)

func newCancelacionAutoMessage(req classes.SaleMessage) {
	FormateaCampo7()
	FormateaCampo11()
	FormateaCampo12()
	FormateaCampo13()
	FormateaCampo17()
	FormateaCampo18()
	FormateaCampo32("12")
}

func VoidMessageContruct(mti string, req request.CancelacionRequest) (classes.ReversalMessage, error) {
	var reversalData classes.ReversalMessage
	respuesta, err := database.GetMessageRequest(req.UUID)
	if err != nil {
		return reversalData, err
	}
	newCancelacionAutoMessage(respuesta)

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
