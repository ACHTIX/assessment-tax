package Tax

import (
	"github.com/ACHTIX/assessment-tax/model"
	"github.com/ACHTIX/assessment-tax/util"
	"testing"
)

func TestTaxCalculation(t *testing.T) {
	// Mock or control the other function calls if necessary
	cases := []struct {
		name          string
		input         model.TaxInput
		expected      float64
		taxLevel      model.TaxLevel
		expectedError string // New field for error message expectation
	}{
		{
			name: "Zero Tax Bracket",
			input: model.TaxInput{
				TotalIncome: 100000,
				Wht:         5000,
				Allowances:  []model.Allowance{},
			},
			expected:      -5000,
			taxLevel:      model.TaxLevel{Level: "0-150,000", Tax: 0},
			expectedError: "", // No error expected for this case
		},
		{
			name: "Negative Total Income",
			input: model.TaxInput{
				TotalIncome: -10000,
				Wht:         2000,
				Allowances:  []model.Allowance{},
			},
			expected:      0,                // Since an error is expected, the result doesn't matter
			taxLevel:      model.TaxLevel{}, // Similarly, tax level doesn't matter in case of error
			expectedError: "\"Key: 'TaxInput.TotalIncome' Error:Field validation for 'TotalIncome' failed on the 'gte' tag\"",
		},
		{
			name: "No Withholding Tax",
			input: model.TaxInput{
				TotalIncome: 50000,
				Wht:         0,
				Allowances:  []model.Allowance{},
			},
			expected:      0,
			taxLevel:      model.TaxLevel{Level: "", Tax: 0}, // Expected tax level for income within 0-150,000 bracket
			expectedError: "",                                // No error expected for this case
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, level, err := util.TaxCalculation(c.input)
			if c.expectedError != "" {
				if err == nil || err.Error() != c.expectedError {
					t.Errorf("Test '%s' failed. Expected error: '%s', Got: '%v'", c.name, c.expectedError, err)
				}
				return // Skip further assertions if error is expected
			}

			if result != c.expected || level != c.taxLevel {
				t.Errorf("Test Failed: %s expected %f and %+v got %f and %+v", c.name, c.expected, c.taxLevel, result, level)
			}
		})
	}
}
