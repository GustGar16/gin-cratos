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
	FieldElement[32] = FieldData{Type: "n", Size: 11, Fixed: true, Mandatory: false, SizePos: "LL", Usage: "Acquiring institution identification code"}
	FieldElement[35] = FieldData{Type: "ans", Size: 37, Fixed: true, Mandatory: false, SizePos: "LL", Usage: "Track 2 Data"}
	FieldElement[37] = FieldData{Type: "an", Size: 12, Fixed: true, Mandatory: true, SizePos: "", Usage: "Retrieval reference number"}
	FieldElement[38] = FieldData{Type: "an", Size: 6, Fixed: true, Mandatory: false, SizePos: "", Usage: "Authorization identification response"}
	FieldElement[39] = FieldData{Type: "an", Size: 2, Fixed: true, Mandatory: false, SizePos: "", Usage: "Response code"}
	FieldElement[41] = FieldData{Type: "ans", Size: 16, Fixed: true, Mandatory: true, SizePos: "", Usage: "Card acceptor terminal identification"}
	FieldElement[42] = FieldData{Type: "ans", Size: 15, Fixed: true, Mandatory: false, SizePos: "", Usage: "Card acceptor identification code"}
	FieldElement[43] = FieldData{Type: "ans", Size: 40, Fixed: true, Mandatory: true, SizePos: "", Usage: "Card acceptor name/location"}
	FieldElement[44] = FieldData{Type: "ans", Size: 4, Fixed: true, Mandatory: false, SizePos: "", Usage: "Additional response data"}
	FieldElement[45] = FieldData{Type: "ans", Size: 76, Fixed: true, Mandatory: false, SizePos: "LL", Usage: "Track 1 Data"}
	FieldElement[48] = FieldData{Type: "ans", Size: 30, Fixed: true, Mandatory: true, SizePos: "LLL", Usage: "Retailer Data"}
	FieldElement[49] = FieldData{Type: "n", Size: 3, Fixed: true, Mandatory: true, SizePos: "LLL", Usage: "Currency code, transaction"}
	FieldElement[54] = FieldData{Type: "ans", Size: 15, Fixed: true, Mandatory: false, SizePos: "LLL", Usage: "Additional amounts"}
	FieldElement[60] = FieldData{Type: "ans", Size: 19, Fixed: true, Mandatory: true, SizePos: "LL", Usage: "Terminal Data"}
	FieldElement[61] = FieldData{Type: "ans", Size: 22, Fixed: true, Mandatory: true, SizePos: "LLL", Usage: "Card Issuer-Category- Response Code Data"}
	FieldElement[62] = FieldData{Type: "ans", Size: 13, Fixed: true, Mandatory: false, SizePos: "", Usage: "Postal Code"}
	FieldElement[63] = FieldData{Type: "ans", Size: 600, Fixed: false, Mandatory: true, SizePos: "LLL", Usage: "Additional Data"}
	FieldElement[70] = FieldData{Type: "n", Size: 3, Fixed: true, Mandatory: false, SizePos: "", Usage: "Network management Information code"}
	FieldElement[90] = FieldData{Type: "an", Size: 42, Fixed: true, Mandatory: false, SizePos: "", Usage: "Original Data Elements"}
	FieldElement[95] = FieldData{Type: "n", Size: 42, Fixed: true, Mandatory: false, SizePos: "", Usage: "Replacement amounts"}
	FieldElement[100] = FieldData{Type: "n", Size: 11, Fixed: false, Mandatory: true, SizePos: "LL", Usage: "Receiving institution identification code"}
	FieldElement[102] = FieldData{Type: "an", Size: 28, Fixed: false, Mandatory: false, SizePos: "LL", Usage: "Account identification 1"}
	FieldElement[120] = FieldData{Type: "ans", Size: 32, Fixed: false, Mandatory: true, SizePos: "LLL", Usage: "Terminal Address-Branch"}
	FieldElement[121] = FieldData{Type: "ans", Size: 23, Fixed: false, Mandatory: true, SizePos: "LLL", Usage: "Authorization Indicators"}
	FieldElement[122] = FieldData{Type: "ans", Size: 14, Fixed: false, Mandatory: false, SizePos: "LLL", Usage: "Card Issuer ID Code"}
	FieldElement[123] = FieldData{Type: "ans", Size: 23, Fixed: false, Mandatory: false, SizePos: "LLL", Usage: "Pos Invoice Data"}
	FieldElement[125] = FieldData{Type: "ans", Size: 15, Fixed: false, Mandatory: true, SizePos: "LLL", Usage: "Pos Settlement Data"}
	FieldElement[126] = FieldData{Type: "ans", Size: 41, Fixed: false, Mandatory: true, SizePos: "LLL", Usage: "Pos Preauthorization and Chrgeback Data"}

}
