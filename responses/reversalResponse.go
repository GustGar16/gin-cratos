package responses

type ReversalResponse struct {
	MTI                 string `json:"mti"`
	ProccesingCode      string `json:"proccesing_code"`
	Amount              string `json:"amount"`
	TransmitionDataTime string `json:"transmition_time"`
	SystemTrace         string `json:"system_trace"`
	LocalTime           string `json:"local_time"`
	LocalDate           string `json:"local_date"`
	MerchantType        string `json:"merchant_type"`
	ServiceEntry        string `json:"service_entry"`
	ServiceCondition    string `json:"service_condition"`
	AcquiringIDCode     string `json:"acquiring_id_code"`
	Track2Data          string `json:"track_two"`
	RetrievalReference  string `json:"retrieval_ref,omitempty"`
	ResponseCode        string `json:"response_code,omitempty"`
	CardAcceptTerminal  string `json:"card_acceptor,omitempty"`
	CurrencyCode        string `json:"currency_code,omitempty"`
	AdditionalAmounts   string `json:"additional_amounts,omitempty"`
	CardIssuer          string `json:"card_issuer,omitempty"`
	PostalCode          string `json:"postal_code,omitempty"`
	AdditionalData      string `json:"additional_data,omitempty"`
	OriginalData        string `json:"original_data,omitempty"`
	AccountID1          string `json:"account_id,omitempty"`
	TerminalAddress     string `json:"terminal_address,omitempty"`
	AuthIndicators      string `json:"auth_indicators,omitempty"`
	PosPreauthData      string `json:"pos_preauth,omitempty"`
}
