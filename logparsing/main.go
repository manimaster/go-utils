package main

import (
	"fmt"
	"logparser"
)

func main() {
	// Define patterns with named capture groups
	patterns := map[string]string{
		"ipPattern":    `(?P<ip>\d+\.\d+\.\d+\.\d+)`,
		"datePattern":  `(?P<date>\d{4}-\d{2}-\d{2})`,
		"timePattern":  `(?P<hour>\d{2}):(?P<minute>\d{2}):(?P<second>\d{2})`,
		"levelPattern": `(?P<level>INFO|ERROR|DEBUG|WARN)`,
	}

	parser := logparser.NewParser(patterns)

	// Sample log line
	logLine := "2023-09-11 12:30:45 INFO 192.168.1.1 Connected successfully"

	results := parser.Parse(logLine)

	// Print results
	for name, matches := range results {
		fmt.Printf("Matches for pattern %s:\n", name)
		for key, value := range matches {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
