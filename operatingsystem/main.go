package main

import (
	"fmt"
	"io/ioutil"
	"osinfo"
)

func main() {
	buf, err := osinfo.GatherSystemInfo(osinfo.CENTOS)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Save buffer to a file
	err = ioutil.WriteFile("system_info.zip", buf.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("System information saved to system_info.zip")
}
