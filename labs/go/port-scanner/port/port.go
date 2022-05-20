package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port string
	State string
}

func ScanPorts(protocols []string, hostname string) []ScanResult {
	var results []ScanResult
	for _, protocol := range(protocols) {
		for port:=0; port<9999; port++ {
			result := ScanResult{Port: protocol + "://" + hostname + ":" + strconv.Itoa(port)}
			address := hostname + ":" + strconv.Itoa(port)
			conn, err := net.DialTimeout(protocol, address, 60*time.Second)

			if err != nil {
				continue
			}
			defer conn.Close()

			result.State = "Open"
			results = append(results, result)
		}

	}
	return results
}

func InitialScan(hostname string) []ScanResult {
	var protocols []string
	protocols = append(protocols, "tcp")
	protocols = append(protocols, "udp")
	results := ScanPorts(protocols, hostname)
	return results
}
