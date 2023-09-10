package main

import (
	"bigfilehandler"
	"encoding/json"
	"fmt"
)

func main() {
	filePath := "path_to_large_file.txt"

	textContent, err := bigfilehandler.ReadFile(filePath, bigfilehandler.TEXT)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File content as text:")
	fmt.Println(textContent)

	jsonContent, err := bigfilehandler.ReadFile(filePath, bigfilehandler.JSON)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonOutput, err := json.MarshalIndent(jsonContent, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File content as JSON:")
	fmt.Println(string(jsonOutput))
}
