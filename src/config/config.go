package config

import (
	promlog "github.com/prometheus/common/log"
)

var log promlog.Logger

func init() {
	log = promlog.Base()
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
