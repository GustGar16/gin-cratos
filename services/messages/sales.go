package messages

import (
	"fmt"
	"gin-Cratos/request"
	"gin-Cratos/responses"
	"strings"
	"time"
)

var campo4, campo7, campo12, campo13, campo17 string

var t = time.Now()

func formateaCampo4(value string) {
	campo4 = strings.Replace(value, ".", "", -1)
	campo4 = fmt.Sprintf("%012s", campo4)

}

func formateaCampo7() {
	campo7 = fmt.Sprintf("%02d%02d%02d%02d%02d",
		t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

}

func formateaCampo12() {
	campo12 = fmt.Sprintf("%02d%02d%02d",
		t.Hour(), t.Minute(), t.Second())
}

func formateaCampo13() {
	campo13 = fmt.Sprintf("%02d%02d",
		t.Month(), t.Day())
}

func formateaCampo17() {
	campo17 = fmt.Sprintf("%02d%02d",
		t.Month(), t.Day())
}

func New(req request.SaleRequest) {
	formateaCampo4(req.Amount)
	formateaCampo7()
	formateaCampo12()
	formateaCampo13()
	formateaCampo17()
}

func SaleMessage(sMti string, req request.SaleRequest) responses.SaleResponse {

	New(req)
	res := responses.SaleResponse{
		MTI:                 sMti,
		ProccesingCode:      req.ProccesingCode,
		Amount:              campo4,
		TransmitionDataTime: campo7,
		SystemTrace:         req.SystemTrace,
		LocalTime:           campo12,
		LocalDate:           campo13,
		DateCapture:         campo17,
		MerchantType:        req.MerchantType,
		ServiceEntry:        req.ServiceEntry,
		ServiceCondition:    req.ServiceCondition,
	}

	return res
}
