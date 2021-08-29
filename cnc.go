package cnc

import (
	"github.com/gin-gonic/gin"
)

type CncServer struct {
	Host     string
	Port     int
	Tor      bool
	Listener *gin.Engine
}

func NewCncServer(host string, port int) CncServer {
	return CncServer{
		Host:     host,
		Port:     port,
		Tor:      false,
		Listener: gin.New(),
	}
}
