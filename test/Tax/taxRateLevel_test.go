package Tax

import (
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestTaxRateLevel(t *testing.T) {
	cases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Negative Income", -100, -1},
		{"Zero Income", 0, 0},
		{"No Tax Bracket", 100000, 0},
		{"Boundary of No Tax and 10%", 150000, 0},
		{"Just Above No Tax Bracket", 150001, 0.10},
		{"Middle of 10% Bracket", 300000, 0.10},
		{"Boundary of 10% and 15%", 500000, 0.10},
		{"Just Above 10% Bracket", 500001, 0.15},
		{"Middle of 15% Bracket", 750000, 0.15},
		{"Boundary of 15% and 20%", 1000000, 0.15},
		{"Just Above 15% Bracket", 1000001, 0.20},
		{"Middle of 20% Bracket", 1500000, 0.20},
		{"Boundary of 20% and 35%", 2000000, 0.20},
		{"Just Above 20% Bracket", 2000001, 0.35},
		{"Above Highest Bracket", 3000000, 0.35},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := util.TaxRateLevel(c.input)
			if result != c.expected {
				t.Errorf("Test Failed: %s expected %f got %f", c.name, c.expected, result)
			}
		})
	}
}
