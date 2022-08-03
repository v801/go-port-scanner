package main

import (
	"fmt"
	port "goportscanner/scans"
)

func main() {
	userPort := getUserInput()
	isZero, isValidPortRange := validateUserInput(userPort)

	if isZero {
		fmt.Println("Select a port between 1-65535.")
		return
	}
	if isValidPortRange {
		fmt.Printf("Scanning TCP Ports 1-%v ...\n", userPort)
		results := port.InitScan("localhost", userPort)
		fmt.Printf("%+v\n", results)
	} else {
		fmt.Println("Select a port between 1-65535.")
		return
	}

}

func getUserInput() (int) {
	var userPort int
	fmt.Println("Enter upper port range:")
	fmt.Scan(&userPort)
	return userPort
}

func validateUserInput(userPort int) (bool, bool) {
	isZero := userPort == 0
	isValidPortRange := userPort >= 1 && userPort <= 65535
	return isZero, isValidPortRange
}
