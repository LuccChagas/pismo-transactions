package models

type Account struct {
	AccountID      uint32 `json:"accountID,omitempty"`
	DocumentNumber string `json:"document_number"`
}
