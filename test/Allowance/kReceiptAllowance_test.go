package Allowance

import (
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestKReceiptAllowance(t *testing.T) {
	cases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Below Cap", 30000, 30000},
		{"At Cap", 50000, 50000},
		{"Above Cap", 55000, 50000},
		{"Zero Receipt", 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := util.KReceiptAllowance(c.input)
			if result != c.expected {
				t.Errorf("Test Failed: %s expected %f got %f", c.name, c.expected, result)
			}
		})
	}
}
