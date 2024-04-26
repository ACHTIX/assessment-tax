package Allowance

import (
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestDonationAllowance(t *testing.T) {
	cases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Below Cap", 50000, 50000},
		{"At Cap", 100000, 100000},
		{"Above Cap", 150000, 100000},
		{"Zero Donation", 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := util.DonationAllowance(c.input)
			if result != c.expected {
				t.Errorf("Test Failed: %s expected %f got %f", c.name, c.expected, result)
			}
		})
	}
}
