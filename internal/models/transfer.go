package models

type Transfer struct {
	Sender string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount float32 `json:"amount"`
}
