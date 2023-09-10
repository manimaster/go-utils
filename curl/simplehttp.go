package simplehttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ResponseJSON is a structure for the JSON response
type ResponseJSON struct {
	StatusCode int                    `json:"status_code"`
	Headers    map[string][]string    `json:"headers"`
	Body       map[string]interface{} `json:"body"`
}

// SendRequest sends an HTTP request and returns the response as a JSON object
func SendRequest(method, url string, headers map[string]string, body []byte) (*ResponseJSON, error) {
	// Create a new request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set the provided headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Print request details
	fmt.Printf("Sending %s request to %s with headers %v and body %s\n", method, url, headers, string(body))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Print response status code
	fmt.Printf("Received response with status code: %d\n", resp.StatusCode)

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(respBody, &jsonResponse)
	if err != nil {
		return nil, err
	}

	return &ResponseJSON{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       jsonResponse,
	}, nil
}
