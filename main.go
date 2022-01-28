package main

import (
	"fmt"

	"github.com/enesffidan/go-portscanner/port"
)

var resultChannel = make(chan bool)

func reciever(resultChannel <-chan bool) bool {
	portStatus := <-resultChannel
	return portStatus
}

func main() {
	fmt.Println("Port Scanning")
	go port.ScanPortInRange("tcp", "google.com", 1, 81, resultChannel)

	portStatus := false
	for i := 1; i <= 81; i++ {
		portStatus = reciever(resultChannel)
		fmt.Println(i, portStatus)
	}

}
