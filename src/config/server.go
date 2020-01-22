package config

import (
	"net"
	"os"
	"strconv"
)

// Server Server configuration
type Server struct {
	listenAddress net.IP
	port          int
}

func setServerConfig() *Server {
	listenAddressString := os.Getenv("SERVER_LISTEN_ADDRESS")
	if listenAddressString == "" {
		listenAddressString = "0.0.0.0"
	}

	listenAddress := net.ParseIP(listenAddressString)
	if listenAddress == nil {
		log.Fatalln("Can't parse listen address")
		return nil
	}

	portString := os.Getenv("SERVER_LISTEN_PORT")
	if portString == "" {
		portString = "8080"
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatalln("Can't parse env SERVER_LISTEN_PORT:", err)
		return nil
	}

	return &Server{
		listenAddress: listenAddress,
		port:          port,
	}
}

// GetListenAddress Return listen address of server configuration
func (s *Server) GetListenAddress() net.IP {
	return s.listenAddress
}

// GetPort Return listen address of server configuration
func (s *Server) GetPort() int {
	return s.port
}
