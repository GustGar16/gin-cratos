package request

type SaleRequest struct {
	MTI                 string `json:"mti"`
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
	AcquiringIDCode     string `json:"acquiring_id_code"`
	RetrievalReference  string `json:"retrieval_ref,omitempty"`
	CardAcceptTerminal  string `json:"card_acceptor,omitempty"`
	CardAcceptorCode    string `json:"card_acceptor_code,omitempty"`
	CardAcceptorName    string `json:"card_acceptor_name,omitempty"`
	RetailerData        string `json:"retailer_data,omitempty"`
	CurrencyCode        string `json:"currency_code,omitempty"`
	AdditionalAmounts   string `json:"additional_amounts,omitempty"`
	TerminalData        string `json:"terminal_data,omitempty"`
	CardIssuer          string `json:"card_issuer,omitempty"`
	PostalCode          string `json:"postal_code,omitempty"`
	AdditionalData      string `json:"additional_data,omitempty"`
	OriginalData        string `json:"original_data,omitempty"`
	ReplacementAmounts  string `json:"replacement_amounts,omitempty"`
	ReceivingIDCode     string `json:"receiving_id_code,omitempty"`
	TerminalAddress     string `json:"terminal_address,omitempty"`
	AuthIndicators      string `json:"auth_indicators,omitempty"`
	PosInvoiceData      string `json:"pos_invoice_data,omitempty"`
	PosSettlementData   string `json:"pos_settlement,omitempty"`
	PosPreauthData      string `json:"pos_preauth,omitempty"`
}
