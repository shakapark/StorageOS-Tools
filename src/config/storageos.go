package config

import "os"

// StorageOSConfig StorageOS API Configuration
type StorageOSConfig struct {
	nodes    string
	username string
	password string
}

func getStorageOSConfig() *StorageOSConfig {
	nodes := os.Getenv("STORAGEOS_ENDPOINT")
	if nodes == "" {
		nodes = "storageos.storageos:5705"
	}

	username := os.Getenv("STORAGEOS_USERNAME")
	if username == "" {
		username = "storageos"
	}

	password := os.Getenv("STORAGEOS_PASSWORD")
	if password == "" {
		log.Fatalln("STORAGEOS_PASSWORD is empty")
		return nil
	}

	return &StorageOSConfig{
		nodes:    nodes,
		username: username,
		password: password,
	}
}

// GetEndpoint Return StorageOS Endpoint
func (sosCfg *StorageOSConfig) GetEndpoint() string {
	return sosCfg.nodes
}

// GetUsername Return StorageOS Username
func (sosCfg *StorageOSConfig) GetUsername() string {
	return sosCfg.username
}

// GetPassword Return StorageOS Password
func (sosCfg *StorageOSConfig) GetPassword() string {
	return sosCfg.password
}
