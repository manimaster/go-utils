package main

import (
	"encoding/json"
	"fmt"
	"simplehttp"
)

func main() {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := []byte(`{"key":"value"}`)

	response, err := simplehttp.SendRequest("POST", "https://jsonplaceholder.typicode.com/posts", headers, body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the response to JSON string for printing
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("JSON Response:")
	fmt.Println(string(responseJSON))
}
