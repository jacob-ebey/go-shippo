package models

// See https://goshippo.com/docs/reference#transactions
type TransactionInput struct {
	Rate          string `json:"rate,omitempty"`
	Metadata      string `json:"metadata,omitempty"`
	LabelFileType string `json:"label_file_type"`
	Async         bool   `json:"async"`

	Shipment         *ShipmentInput `json:"shipment,omitempty"`              // instant call only: https://goshippo.com/docs/reference#transactions-create-instant
	CarrierAccount   string         `json:"carrier_account,omitempty"`       // instant call only: https://goshippo.com/docs/reference#transactions-create-instant
	ServerLevelToken string         `json:"servericelevel_token, omitempty"` // instant call only: https://goshippo.com/docs/reference#transactions-create-instant
}

// See https://goshippo.com/docs/reference#transactions
type Transaction struct {
	TransactionInput
	CommonOutputFields
	ObjectStatus         string                `json:"object_status,omitempty"`
	Status               string                `json:"status,omitempty"`
	Test                 bool                  `json:"test"`
	TrackingNumber       string                `json:"tracking_number,omitempty"`
	TrackingStatus       *TrackingStatusDict   `json:"tracking_status,omitempty"`
	TrackingHistory      []*TrackingStatusDict `json:"tracking_history,omitempty"`
	TrackingURLProvider  string                `json:"tracking_url_provider,omitempty"`
	LabelURL             string                `json:"label_url,omitempty"`
	CommercialInvoiceURL string                `json:"commercial_invoice_url,omitempty"`
	Messages             []*OutputMessage      `json:"messages,omitempty"`
	Async                bool                  `json:"async"`
}
