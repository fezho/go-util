package util

import (
	"net"
	"os"
)

// GetLocalIP check the value of POD_IP first, or returns the non loopback local IP of the host
func GetLocalIP() string {
	ip := os.Getenv("POD_IP")
	if len(ip) != 0 {
		return ip
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback then display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

