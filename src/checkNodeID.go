package main

import "github.com/shakapark/storageos-tools/src/storageos"

func checkNodeID(etcdURLs []string, hostname, filePath string) bool {
	log.Infoln("Check Node ID")

	clientETCD, err := storageos.NewClientETCD(etcdURLs, "", "")
	if err != nil {
		log.Fatalln("Errors client ETCD: ", err)
	}

	defer clientETCD.Close()

	nodeIDetcd, err := storageos.GetETCDNodeID(clientETCD, hostname)
	if err != nil {
		log.Errorln("Errors NodeID from Etcd: ", err)
		return false
	}
	log.Infoln("NodeID of", hostname, ":", nodeIDetcd)

	nodeIDfile, err := storageos.GetFileID(filePath)
	if err != nil {
		log.Errorln("Errors NodeID from File: ", err)
		return false
	}
	log.Infoln("NodeID of file", nodeIDfile)

	if nodeIDetcd != nodeIDfile {
		log.Infoln("Ids different, change it...")
		err := storageos.ReplaceFileID(filePath, nodeIDetcd)
		if err != nil {
			log.Errorln("Errors Writing new ID: ", err)
			return false
		}
	} else {
		log.Infoln("Same Ids")
	}

	return true
}
