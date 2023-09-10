package netinfo

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/net"
)

// DataType represents the type of data to return
type DataType int

const (
	// STRING returns the data as plain string
	STRING DataType = iota
	// JSONDATA returns the data as a JSON object
	JSONDATA
)

// GetNetworkInfo retrieves information about the specified network adapters
func GetNetworkInfo(adapters []string, dataType DataType) (interface{}, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get network interfaces: %w", err)
	}

	var filteredInterfaces []net.InterfaceStat

	for _, intf := range interfaces {
		for _, adapter := range adapters {
			if intf.Name == adapter {
				filteredInterfaces = append(filteredInterfaces, intf)
				break
			}
		}
	}

	switch dataType {
	case STRING:
		return fmt.Sprint(filteredInterfaces), nil

	case JSONDATA:
		data, err := json.Marshal(filteredInterfaces)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
		}
		return string(data), nil

	default:
		return nil, fmt.Errorf("unsupported data type")
	}
}
