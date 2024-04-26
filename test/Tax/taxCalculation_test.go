package Tax

import (
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestTaxCalculation(t *testing.T) {
	// Mock or control the other function calls if necessary
	cases := []struct {
		name     string
		input    model.TaxInput
		expected float64
		taxLevel model.TaxLevel
	}{
		{
			name: "Zero Tax Bracket",
			input: model.TaxInput{
				TotalIncome: 100000,
				Wht:         5000,
				Allowances:  []model.Allowance{},
			},
			expected: -5000,
			taxLevel: model.TaxLevel{Level: "0-150,000", Tax: 0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, level := util.TaxCalculation(c.input)
			if result != c.expected || level != c.taxLevel {
				t.Errorf("Test Failed: %s expected %f and %+v got %f and %+v", c.name, c.expected, c.taxLevel, result, level)
			}
		})
	}
}
