package main

import (
	"fmt"
	"netinfo"
)

func main() {
	// Get information about "eth0" and "lo" interfaces
	data, err := netinfo.GetNetworkInfo([]string{"eth0", "lo"}, netinfo.JSONDATA)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Network Information:", data)
}
