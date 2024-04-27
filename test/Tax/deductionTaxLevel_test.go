package Tax

import (
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestDeductionTaxLevel(t *testing.T) {
	cases := []struct {
		name     string
		input    float64
		expected float64
		taxLevel model.TaxLevel
	}{
		{
			name:     "Middle Bracket",
			input:    750000,
			expected: 72500, // Computed tax
			taxLevel: model.TaxLevel{Level: "500,001-1,000,000", Tax: 72500},
		},
		{
			name:     "Lower Bracket",
			input:    10000,
			expected: 0, // No tax for netIncome ≤ 150,000
			taxLevel: model.TaxLevel{Level: "0-150,000", Tax: 0.0},
		},
		{
			name:     "Upper Bracket",
			input:    2500000,
			expected: 485000, // Computed tax for netIncome >= 2,000,001
			taxLevel: model.TaxLevel{Level: "2,000,001 ขึ้นไป", Tax: 485000},
		},
		{
			name:     "Boundary Case - Just Below Bracket",
			input:    150000,
			expected: 0, // No tax for netIncome ≤ 150,000
			taxLevel: model.TaxLevel{Level: "0-150,000", Tax: 0.0},
		},
		{
			name:     "Boundary Case - Just Above Bracket",
			input:    500000,
			expected: 35000, // Computed tax for netIncome between 150,001 and 500,000
			taxLevel: model.TaxLevel{Level: "150,001-500,000", Tax: 35000},
		},
		{
			name:     "Negative Net Income",
			input:    -1000,
			expected: -1,               // Error case for negative netIncome
			taxLevel: model.TaxLevel{}, // No tax level expected for error case
		},
		{
			name:     "Zero Net Income",
			input:    0,
			expected: 0, // No tax for netIncome = 0
			taxLevel: model.TaxLevel{Level: "0-150,000", Tax: 0.0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, level := util.DeductionTaxLevel(c.input)
			if result != c.expected || level != c.taxLevel {
				t.Errorf("Test Failed: %s expected %f and %+v got %f and %+v", c.name, c.expected, c.taxLevel, result, level)
			}
		})
	}
}
