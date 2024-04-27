package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// TestTaxCalculationsUploadCsv handles JSON response comparison more effectively.
func TestTaxCalculationsUploadCsv(t *testing.T) {
	url := "http://localhost:8080/tax/calculations/upload-csv"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open("./tax_data.csv")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("taxFile", filepath.Base("./tax_data.csv"))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy data to form file: %v", err)
	}

	if err = writer.Close(); err != nil {
		t.Fatalf("Failed to close writer: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var actualResponse interface{}
	var expectedResponse interface{}
	expectedResponseBody := `{
    "taxes": [
        {
            "totalIncome": 500000,
            "tax": 29000
        },
        {
            "totalIncome": 600000,
            "taxRefund": 2000
        },
        {
            "totalIncome": 750000,
            "tax": 11250
        }
    ]
}`

	json.Unmarshal([]byte(expectedResponseBody), &expectedResponse)
	json.Unmarshal(body, &actualResponse)

	if !reflect.DeepEqual(actualResponse, expectedResponse) {
		t.Errorf("Unexpected response body: got %v want %v", string(body), expectedResponseBody)
	}
}
