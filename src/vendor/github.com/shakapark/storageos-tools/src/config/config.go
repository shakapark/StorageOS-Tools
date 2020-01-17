package config

import (
	"net"
	"os"
	"strconv"
	"strings"

	promlog "github.com/prometheus/common/log"
)

var log promlog.Logger

func init() {
	log = promlog.Base()
}

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

// CheckID CheckID configuration
type CheckID struct {
	etcdURLs []string
	hostname string
	pathFile string
}

func setcheckIDConfig() *CheckID {
	etcdURLsString := os.Getenv("ETCD_URLS")
	if etcdURLsString == "" {
		log.Fatalln("ETCD_URLS is empty")
	}
	etcdURLsStrings := strings.Split(etcdURLsString, ",")

	hostname := os.Getenv("NODE_HOSTNAME")
	if hostname == "" {
		log.Fatalln("NODE_HOSTNAME is empty")
		return nil
	}

	pathFile := os.Getenv("NODE_PATHFILE")
	if pathFile == "" {
		log.Fatalln("NODE_PATHFILE is empty")
		return nil
	}

	return &CheckID{
		etcdURLs: etcdURLsStrings,
		hostname: hostname,
		pathFile: pathFile,
	}

}

// GetEtcdURLs Return ETCD Urls of CheckID configuration
func (cid *CheckID) GetEtcdURLs() []string {
	return cid.etcdURLs
}

// GetHostname Return ETCD Urls of CheckID configuration
func (cid *CheckID) GetHostname() string {
	return cid.hostname
}

// GetPathFile Return ETCD Urls of CheckID configuration
func (cid *CheckID) GetPathFile() string {
	return cid.pathFile
}

// Config Contains the server Config
type Config struct {
	server  *Server
	checkID *CheckID
}

// GetConfig Return new configuration
func GetConfig() *Config {
	server := setServerConfig()
	if server == nil {
		return nil
	}

	checkID := setcheckIDConfig()
	if checkID == nil {
		return nil
	}

	return &Config{
		server:  server,
		checkID: checkID,
	}
}

// GetServerConfig Return server config of Config
func (c *Config) GetServerConfig() *Server {
	return c.server
}

// GetCheckIDConfig Return checkID config of Config
func (c *Config) GetCheckIDConfig() *CheckID {
	return c.checkID
}
