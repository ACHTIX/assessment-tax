package Allowance

import (
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestAllowanceKReceiptAdmin(t *testing.T) {
	cases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Zero", 0, 0},
		{"Normal Case", 50000, 50000},
		{"Just Above Cap", 100000, 100000},
		{"Much Above Cap", 200000, 100000},
		{"At Cap", 100000, 100000},
		{"Minus", -100000, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := util.AllowanceKReceiptAdmin(c.input)
			if result != c.expected {
				t.Errorf("Test Failed: %s expected %f got %f", c.name, c.expected, result)
			}
		})
	}
}
