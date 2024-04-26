package Tax

import (
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestIncomeCalculation(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		totalIncome    float64
		allowances     []model.Allowance
		expectedIncome float64
	}{
		{
			name:           "No additional allowances",
			totalIncome:    100000,
			allowances:     []model.Allowance{}, // No additional allowances
			expectedIncome: 40000,               // 100000 - 60000
		},
		{
			name:        "With additional allowances",
			totalIncome: 150000,
			allowances: []model.Allowance{
				{AllowanceType: "donation", Amount: 30000},
			},
			expectedIncome: 60000, // 150000 - (60000 + 30000)
		},
		{
			name:           "Zero income",
			totalIncome:    0,
			allowances:     []model.Allowance{},
			expectedIncome: -60000, // 0 - 60000
		},
	}

	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Mock the CheckAllowanceType function if necessary or handle it according to actual implementation
			actualIncome := util.IncomeCalculation(model.TaxInput{TotalIncome: tc.totalIncome, Allowances: tc.allowances})
			if actualIncome != tc.expectedIncome {
				t.Errorf("Failed %s: expected income %v, got %v", tc.name, tc.expectedIncome, actualIncome)
			}
		})
	}
}
