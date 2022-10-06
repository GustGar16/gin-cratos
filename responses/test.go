package responses

type TestResponse struct {
	ProccesingCode      string `json:"proccesing_code"`
	Amount              string `json:"amount"`
	TransmitionDataTime string `json:"transmition_time"`
	SystemTrace         string `json:"system_trace"`
	LocalTime           string `json:"local_time"`
	LocalDate           string `json:"local_date"`
	DateCapture         string `json:"date_capture"`
	MerchantType        string `json:"merchant_type"`
	ServiceEntry        string `json:"service_entry"`
	ServiceCondition    string `json:"service_condition"`
}
