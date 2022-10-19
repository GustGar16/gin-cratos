package struct

import (
	"fmt"
	"gin-Cratos/models/fields"
	"gin-Cratos/models/header"
)

/*headersConstruct headers data */
func headersConstruct() {
	headerMti := header.HeaderData{Type: "n", Size: 4, Fixed: true, Mandatory: true, Usage: "MTI"}
	headerBitmap := header.HeaderData{Type: "n", Size: 8, Fixed: true, Mandatory: true, Usage: "Bit Map Primary"}
	headers := header.HeaderElement{MTI: headerMti, BITMAP: headerBitmap}
	fmt.Println(headers)
}

/*fieldsConstruct Fields data */
func fieldsConstruct() {
	fields.NewFieldElement()
	fmt.Println(fields.FieldElement)
}

/*New constructor de la estructura de campos*/
func New() {

	headersConstruct()

	fieldsConstruct()

}
