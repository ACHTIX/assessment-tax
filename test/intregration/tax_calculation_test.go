package integration

import (
	"encoding/json"
	"fmt"
	"github.com/ACHTIX/assessment-tax/model"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// TestTaxCalculationFunction tests the POST request handling of the tax calculation API.
func TestTaxCalculationFunction(t *testing.T) {

	url := `http://localhost:8080/tax/calculations`
	method := "POST"

	payload := strings.NewReader(`{
        "totalIncome": 500000.0,
        "wht": 500000,
        "allowances": [
            {
                "allowanceType": "k-receipt",
                "amount": 100000.0
            },
            {
                "allowanceType": "donation",
                "amount": 100000.0
            }
        ]
    }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var actualResponse model.TaxData
	if err := json.Unmarshal(body, &actualResponse); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	expectedResponse := model.TaxData{
		TaxRefund: 486000,
		TaxLevel: []model.TaxLevel{
			{"0-150,000", 0},
			{"150,001-500,000", 14000},
			{"500,001-1,000,000", 0},
			{"1,000,001-2,000,000", 0},
			{"2,000,001 ขึ้นไป", 0},
		},
	}
	if !compareTaxResponses(expectedResponse, actualResponse) {
		t.Errorf("Unexpected response: got %+v want %+v", actualResponse, expectedResponse)
	}

}

func compareTaxResponses(a, b model.TaxData) bool {
	if a.Tax != b.Tax || a.TaxRefund != b.TaxRefund || len(a.TaxLevel) != len(b.TaxLevel) {
		return false
	}
	for i := range a.TaxLevel {
		if a.TaxLevel[i].Level != b.TaxLevel[i].Level || a.TaxLevel[i].Tax != b.TaxLevel[i].Tax {
			fmt.Printf("Mismatch found in TaxLevel[%d]: Expected %v, got %v\n", i, a.TaxLevel[i], b.TaxLevel[i])
			return false
		}
	}
	return true
}
