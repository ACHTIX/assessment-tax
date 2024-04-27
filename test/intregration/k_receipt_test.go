package integration

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// Function to test
func sendPostRequest(url string, input string) (string, error) {
	method := "POST"
	payload := strings.NewReader(input) // Use the input parameter for the payload
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic YWRtaW5UYXg6YWRtaW4h")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil // Trim space and newline characters
}

// Unit test
func TestSendPostRequest(t *testing.T) {
	// Run the function with the URL of the mock server
	body, err := sendPostRequest("http://localhost:8080/admin/deductions/k-receipt", `{"amount": 0.1}`)
	if err != nil {
		t.Fatal(err)
	}

	expectedBody := `{"kReceipt":0.1}`
	if body != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, body)
	}
}
