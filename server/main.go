package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/martins0n/openvpn-traffic-viewer/lib"
)

func ShowStats(logPath string) *[]lib.OpenVpnStatus {
	openVpnStatusList := lib.ParseOpenVpnStatus(logPath)
	return openVpnStatusList
}

func main() {
	var logPath string
	var port string
	var address string

	flag.StringVar(&logPath, "log", "", "log file path")
	flag.StringVar(&port, "port", "8080", "port")
	flag.StringVar(&address, "address", "0.0.0.0", "address")

	flag.Parse()

	r := gin.Default()
	r.GET("/showstats", func(c *gin.Context) {
		c.JSON(200, ShowStats(logPath))
	})
	r.Run(address + ":" + port)
}
