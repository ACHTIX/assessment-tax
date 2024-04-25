package model

type TaxResult struct {
	Tax float64 `json:"tax"`
}

type TaxData struct {
	Tax       float64    `json:"tax,omitempty"`
	TaxRefund float64    `json:"taxRefund,omitempty"`
	TaxLevel  []TaxLevel `json:"taxLevel"`
}

type TaxLevel struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

type AdminResponseStruct struct {
	PersonalDeduction float64 `json:"personalDeduction,omitempty" validate:"gte=0"`
	KReceipt          float64 `json:"kReceipt,omitempty" validate:"gte=0"`
	Donation          float64 `json:"donation,omitempty" validate:"gte=0"`
}
