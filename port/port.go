package port

import (
	"net"
	"strconv"
	"time"

	"honnef.co/go/netdb"
)

type ScanResult struct {
	Port    int
	Status  string
	Service string
}

func ScanPort(protocol string, hostname string, port int) ScanResult { //connection scan to single port
	address := hostname + ":" + strconv.Itoa(port)
	portScanResult := ScanResult{port, "", ""}

	var protoent *netdb.Protoent
	var servent *netdb.Servent

	conn, err := net.DialTimeout(protocol, address, 5*time.Second)

	if err != nil {
		portScanResult.Status = "close"
		return portScanResult
	} else {
		portScanResult.Status = "open"
	}

	protoent = netdb.GetProtoByName(protocol)
	servent = netdb.GetServByPort(port, protoent)

	portScanResult.Service = servent.Name

	defer conn.Close()

	return portScanResult

}

func ScanPortInRange(protocol string, hostname string, start int, end int) map[int]ScanResult { //connection scan ports of given range

	resultMap := make(map[int]ScanResult)
	for i := start; i <= end; i++ {

		scanMessage := ScanPort(protocol, hostname, i)
		resultMap[i] = scanMessage

	}

	return resultMap

}

// func PortScanResult(protocol string, hostname string, port int, status bool) string {

// 	var protoent *netdb.Protoent
// 	var servent *netdb.Servent

// 	protoent = netdb.GetProtoByName(protocol)
// 	servent = netdb.GetServByPort(port, protoent)

// 	var portStatus string
// 	if status == true {
// 		portStatus = "open"
// 	} else {
// 		portStatus = "close"
// 	}
// 	message := fmt.Sprintf("Port:%d\\%s Status:%s Service:%s", port, protocol, portStatus, servent.Name)
// 	return message
// }
