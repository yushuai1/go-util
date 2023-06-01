package nets

import (
	"net"
	"os"
	"strings"
	"sync"
)

var myip string
var once sync.Once

func GetIP() string {

	once.Do(func() {
		conn, _ := net.Dial("udp", "8.8.8.8:80")
		defer conn.Close()
		localAddr := conn.LocalAddr().String()
		idx := strings.LastIndex(localAddr, ":")
		myip = localAddr[0:idx]
	})
	return myip
}
func GetMacAddr() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	if len(interfaces) == 0 {
		return ""
	}

	maxIndexInterface := interfaces[0]
	for _, inter := range interfaces {
		if inter.HardwareAddr == nil {
			continue
		}
		if inter.Flags&net.FlagUp == 1 {
			maxIndexInterface = inter
		}
	}
	return maxIndexInterface.HardwareAddr.String()
}

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return ""
	}
	return name
}
