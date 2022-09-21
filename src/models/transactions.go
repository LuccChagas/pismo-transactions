package models

type Transaction struct {
	TransactionID   uint32  `json:"transaction_id,omitempty"`
	AccountID       uint32  `json:"account_id,omitempty"`
	OperationTypeID uint32  `json:"operation_type_id,omitempty"`
	Amount          float32 `json:"amount,omitempty"`
	EventDate       string  `json:"event_date,omitempty"`
}
