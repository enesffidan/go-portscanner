package port

import (
	"fmt"
	"net"
	"ping"
	"strconv"
	"time"
)

func ScanPort(protocol string, hostname string, port int) bool { //connection scan to single port
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 2*time.Second)

	if err != nil {
		fmt.Printf("ERROR: %s \n", err.Error())
		return false
	}
	defer conn.Close()
	return true

}

func ScanPortInRange(protocol string, hostname string, start int, end int, resultChannel chan bool) { //connection scan ports of given range

	//var results []bool
	for i := start; i <= end; i++ {
		resultChannel <- ScanPort(protocol, hostname, i)
		//results = append(results, ScanPort(protocol, hostname, i))
	}

	//return results
}

func PingPort(hostname string, count int, port int) {
	pinger, err := ping.NewPinger(hostname)

	if err != nil {
		fmt.Printf("ERROR: %s \n", err.Error())
	}

	pinger.Count = count
	pinger.Run()

}
