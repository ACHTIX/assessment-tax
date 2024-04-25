package model

type TaxResponseCSVDataStruct struct {
	TotalIncome float64 `json:"totalIncome" validate:"gte=0"`
	Tax         float64 `json:"tax,omitempty" validate:"gte=0"`
	TaxRefund   float64 `json:"taxRefund,omitempty" validate:"gte=0"`
}

type TaxResponseCSVStruct struct {
	Taxes []TaxResponseCSVDataStruct `json:"taxes" validate:"gte=0"`
}
