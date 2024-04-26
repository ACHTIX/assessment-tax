package Allowance

import (
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestAllowancePersonalAdmin(t *testing.T) {
	testCases := []struct {
		name   string
		amount float64
		want   float64
	}{
		{"Exactly 100001", 100001, 100000},
		{"Above 100000", 150000, 100000},
		{"Exactly 100000", 100000, 100000},
		{"Less than 10000", 5000, 10000},
		{"Exactly 10000", 10000, 10000},
		{"Within valid range", 50000, 50000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := util.AllowancePersonalAdmin(tc.amount)
			if got != tc.want {
				t.Errorf("AllowancePersonalAdmin(%f) = %f; want %f", tc.amount, got, tc.want)
			}
		})
	}
}
