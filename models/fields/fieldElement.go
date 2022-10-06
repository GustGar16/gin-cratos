package fields

var FieldElement = make(map[int]FieldData)

func NewFieldElement() {
	FieldElement[1] = FieldData{Type: "b", Size: 16, Fixed: true, Mandatory: true, SizePos: "", Usage: "Bit Map Extended"}
	FieldElement[3] = FieldData{Type: "n", Size: 6, Fixed: true, Mandatory: true, SizePos: "", Usage: "Processing code"}
	FieldElement[4] = FieldData{Type: "n", Size: 12, Fixed: true, Mandatory: false, SizePos: "", Usage: "Amount, transaction"}
	FieldElement[7] = FieldData{Type: "n", Size: 10, Fixed: true, Mandatory: false, SizePos: "", Usage: "Transmission, date & time"}
	FieldElement[11] = FieldData{Type: "n", Size: 6, Fixed: true, Mandatory: false, SizePos: "", Usage: "System trace audit number"}
	FieldElement[12] = FieldData{Type: "n", Size: 6, Fixed: true, Mandatory: false, SizePos: "", Usage: "Time, local transaction"}
	FieldElement[13] = FieldData{Type: "n", Size: 4, Fixed: true, Mandatory: false, SizePos: "", Usage: "Date, local transaction (MMdd)"}
	FieldElement[17] = FieldData{Type: "n", Size: 4, Fixed: true, Mandatory: false, SizePos: "", Usage: "Date, capture"}
	FieldElement[18] = FieldData{Type: "n", Size: 4, Fixed: true, Mandatory: false, SizePos: "", Usage: "Merchant type"}
	FieldElement[22] = FieldData{Type: "n", Size: 3, Fixed: true, Mandatory: false, SizePos: "", Usage: "Point of service entry mode"}
	FieldElement[25] = FieldData{Type: "n", Size: 3, Fixed: true, Mandatory: false, SizePos: "", Usage: "Point of service condition mode"}
}
