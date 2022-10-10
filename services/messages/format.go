package messages

import (
	"fmt"
	"strings"
	"time"

	"gin-Cratos/helpers"
)

var t = time.Now()

var Campo3, Campo4, Campo7, Campo11, Campo12, Campo13, Campo17, Campo18, Campo22, Campo25, Campo32, Campo35, Campo37 string
var Campo41 = "0000CP01        "
var Campo42 string //Pendiente de revisar
var Campo43 string //Pendiente de revisar
var Campo48, Campo49 string
var Campo60 = "CLPGTES1+0000000"

func FormateaCampo3(tipo string) {
	result := ""
	if tipo == "" {
		result += "00"
	} else if tipo == "compra" {
		result += "00"
	} else if tipo == "devolucion" {
		result += "20"
	}
	//Revisar tipo de cuenta (default 00)
	result += "00"
	//Revisar cuenta destino (default 00)
	result += "00"

	Campo3 = result
}

func FormateaCampo4(value string) {
	data := strings.Replace(value, ".", "", -1)
	Campo4 = fmt.Sprintf("%012s", data)

}

func FormateaCampo7() {
	Campo7 = fmt.Sprintf("%02d%02d%02d%02d%02d",
		t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func FormateaCampo11() {
	Campo11 = helpers.NewRandomNumber(6)
}

func FormateaCampo12() {
	Campo12 = fmt.Sprintf("%02d%02d%02d",
		t.Hour(), t.Minute(), t.Second())
}

func FormateaCampo13() {
	Campo13 = fmt.Sprintf("%02d%02d",
		t.Month(), t.Day())
}

func FormateaCampo17() {
	Campo17 = fmt.Sprintf("%02d%02d",
		t.Month(), t.Day())
}

// TODO: revisar el giro de comercio para verificar el return
func FormateaCampo18() {
	Campo18 = "1234"
}

func FormateaCampo22(tipo string) {
	result := ""
	switch tipo {
	case "moto":
	case "devolucion_moto":
		result = "010"
	default:
		result = "012"
	}
	Campo22 = result
}

// TODO: Revisar si este Campo es obligatorio o solo usara el default (00)
func FormateaCampo25() {
	Campo25 = "00"
}

func FormateaCampo32(value string) {
	Campo32 = fmt.Sprintf("%011s", value)
}

func FormateaCampo35(pan, anio, mes string) {
	Campo35 += pan
	Campo35 += "="
	Campo35 += fmt.Sprintf("%02s", anio)
	Campo35 += fmt.Sprintf("%02s", mes)
}

func FormateaCampo37() {
	Campo37 = helpers.NewRandomNumber(12)
}

func FormateaCampo48(afiliacion string) {
	Campo48 += afiliacion
	Campo48 += "            00000000"
}

func FormateaCampo49(moneda string) {
	if moneda == "pesos" {
		Campo49 = "484"
	} else {
		Campo49 = "840"
	}
}
