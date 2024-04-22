# Variable , Function and Definition

## Struct

```
type TaxInput struct {
TotalIncome float64     `json:"totalIncome" validate:"gte=0"`
Wht         float64     `json:"wht" validate:"gte=0"`
Allowances  []Allowance `json:"allowances" validate:"dive"`
}
```

```
type Allowance struct {
AllowanceType string  `json:"allowanceType" validate:"oneof=donation k-receipt"`
Amount        float64 `json:"amount" validate:"gte=0"`
}
```

```
type TaxResult struct {
	Tax float64 `json:"tax"`
}
```

## Function

### checkTaxLevel

### multiplyTaxLevel

### subTaxLevel

### TaxCalculation