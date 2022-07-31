package main

import (
	"fmt"
	port "goportscanner/scans"
)

func main() {
	var userPort int
	fmt.Println("Enter upper port range:")
	fmt.Scan(&userPort)
	fmt.Println("Scanning Ports...")
	results := port.InitScan("localhost", userPort)
	fmt.Printf("%+v\n", results)
}
