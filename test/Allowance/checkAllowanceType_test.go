package Allowance

import (
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestCheckAllowanceType(t *testing.T) {
	cases := []struct {
		name     string
		input    []model.Allowance
		expected float64
	}{
		{"No Allowances", nil, 0},
		{"Only Donations", []model.Allowance{{AllowanceType: "donation", Amount: 150000}, {AllowanceType: "donation", Amount: 50000}}, 150000},
		{"Only K-Receipts", []model.Allowance{{AllowanceType: "k-receipt", Amount: 60000}, {AllowanceType: "k-receipt", Amount: 25000}}, 75000},
		{"Mixed Allowances", []model.Allowance{{AllowanceType: "donation", Amount: 100000}, {AllowanceType: "k-receipt", Amount: 50000}}, 150000},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := util.CheckAllowanceType(c.input)
			if result != c.expected {
				t.Errorf("Test Failed: %s expected %f got %f", c.name, c.expected, result)
			}
		})
	}
}
