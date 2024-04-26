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
