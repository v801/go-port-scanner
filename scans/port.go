package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  int
	State string
}

func scanPort(protocol string, hostname string, port int) ScanResult {
	result := ScanResult{Port: port}
	// convert port int into string and create full address string
	address := hostname + ":" + strconv.Itoa(port)
	// get connection or error from dial attempt with 1min timeout
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	// if we cant connect, set closed state and return result
	if err != nil {
		result.State = "Closed"
		return result
	}

	// when we connect, set open state and return result, close connection once returned
	defer conn.Close()
	result.State = "Open"
	return result
}

// function is Public so we can import into our main go file
func InitScan(hostname string, port int) []ScanResult {
	var results []ScanResult
	for i := 0; i <= port; i++ {
		// loop thru upper port range input by user
		portInfo := scanPort(("tcp"), hostname, i)
		openPort := portInfo.State == "Open"
		// if state is open append results
		if openPort {
			results = append(results, portInfo)
		}
	}
	return results
}
