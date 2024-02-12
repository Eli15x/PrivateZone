package model

type Sale struct {
	SaleId             string    `json:"saleId,omitempty" bson:"saleId,omitempty"`
	UserId             string    `json:"userId,omitempty" bson:"UserId,omitempty"`
	Value              string   `json:"value,omitempty" bson:"value,omitempty"`
	TypePayment        string   `json:"typePayment,omitempty" bson:"typePayment,omitempty"`
	Bank        string   `json:"typePayment,omitempty" bson:"typePayment,omitempty"`
}
