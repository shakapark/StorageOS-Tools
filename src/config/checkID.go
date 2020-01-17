package config

import (
	"os"
	"strings"
)

// CheckID CheckID configuration
type CheckID struct {
	etcdURLs []string
	hostname string
	// pathFile string
	storageos *StorageOSConfig
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

	// pathFile := os.Getenv("NODE_PATHFILE")
	// if pathFile == "" {
	// 	log.Fatalln("NODE_PATHFILE is empty")
	// 	return nil
	// }

	storageosCfg := getStorageOSConfig()
	if storageosCfg == nil {
		return nil
	}

	return &CheckID{
		etcdURLs: etcdURLsStrings,
		hostname: hostname,
		// pathFile: pathFile,
		storageos: storageosCfg,
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
// func (cid *CheckID) GetPathFile() string {
// 	return cid.pathFile
// }

// GetStorageOSConf Return ETCD Urls of CheckID configuration
func (cid *CheckID) GetStorageOSConf() *StorageOSConfig {
	return cid.storageos
}
