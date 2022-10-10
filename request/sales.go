package request

type SaleRequest struct {
	ProccesingCode     string `json:"proccesing_code,omitempty"`
	Amount             string `json:"amount,omitempty"`
	SystemTrace        string `json:"system_trace,omitempty"`
	MerchantType       string `json:"merchant_type,omitempty"`
	ServiceEntry       string `json:"service_entry,omitempty"`
	ServiceCondition   string `json:"service_condition,omitempty"`
	AcquiringIDCode    string `json:"acquiring_id_code,omitempty"`
	RetrievalReference string `json:"retrieval_ref,omitempty"`
	CardAcceptTerminal string `json:"card_acceptor,omitempty"`
	CardAcceptorCode   string `json:"card_acceptor_code,omitempty"`
	CardAcceptorName   string `json:"card_acceptor_name,omitempty"`
	RetailerData       string `json:"retailer_data,omitempty"`
	CurrencyCode       string `json:"currency_code,omitempty"`
	AdditionalAmounts  string `json:"additional_amounts,omitempty"`
	TerminalData       string `json:"terminal_data,omitempty"`
	CardIssuer         string `json:"card_issuer,omitempty"`
	PostalCode         string `json:"postal_code,omitempty"`
	AdditionalData     string `json:"additional_data,omitempty"`
	OriginalData       string `json:"original_data,omitempty"`
	ReplacementAmounts string `json:"replacement_amounts,omitempty"`
	ReceivingIDCode    string `json:"receiving_id_code,omitempty"`
	TerminalAddress    string `json:"terminal_address,omitempty"`
	AuthIndicators     string `json:"auth_indicators,omitempty"`
	PosInvoiceData     string `json:"pos_invoice,omitempty"`
	PosSettlementData  string `json:"pos_settlement,omitempty"`
	PosPreauthData     string `json:"pos_preauth,omitempty"`
}
