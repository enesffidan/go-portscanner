package main

import (
	"fmt"

	"github.com/enesffidan/go-portscanner/port"
)

func main() {
	fmt.Println("Port Scanning")

	//port.PingHost("192.168.150.129", 3)

	//portScanResult, _ := port.ScanPort("tcp", "google.com", 80)

	//fmt.Println(portScanResult)

	m := make(map[int]port.ScanResult)
	m = port.ScanPortInRange("tcp", "google.com", 75, 81)
	fmt.Println(m)

	//port.ScanPortInRange("tcp", "google.com", 1, 10)

}
