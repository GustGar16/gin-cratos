package request

type TestRequest struct {
	ProccesingCode   string `json:"proccesing_code,omitempty"`
	Amount           string `json:"amount,omitempty"`
	SystemTrace      string `json:"system_trace,omitempty"`
	MerchantType     string `json:"merchant_type,omitempty"`
	ServiceEntry     string `json:"service_entry,omitempty"`
	ServiceCondition string `json:"service_condition,omitempty"`
}
