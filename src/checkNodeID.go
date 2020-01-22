package main

import "StorageOS-Tools/src/storageos"

func checkNodeID(etcdURLs []string, hostname, filePath, storageosEndpoint, storageosUsername, storageosPassword string) bool {
	log.Infoln("Check Node ID")

	clientETCD, err := storageos.NewClientETCD(etcdURLs, "", "")
	if err != nil {
		log.Fatalln("Errors client ETCD: ", err)
	}

	defer clientETCD.Close()

	nodeIDetcd, err := storageos.GetETCDNodeID(clientETCD, hostname)
	if err != nil {
		log.Errorln("Errors NodeID from Etcd:", err)
		return false
	}
	log.Infoln("NodeID of", hostname, "in etcd:", nodeIDetcd)

	nodeIDfile, err := storageos.GetFileID(filePath)
	if err != nil {
		log.Errorln("Errors NodeID from File:", err)
		return false
	}
	log.Infoln("NodeID of", hostname, "in file:", nodeIDfile)

	if nodeIDetcd != nodeIDfile {
		log.Infoln("Ids different, change it...")
		// err := storageos.ReplaceFileID(filePath, nodeIDetcd)
		err := storageos.DeleteStorageOSNode(nodeIDetcd, storageosEndpoint, storageosUsername, storageosPassword)
		// err := storageos.ListeStorageOSNode(hostname, storageosEndpoint, storageosUsername, storageosPassword)
		if err != nil {
			log.Errorln("Errors Delete old ID:", err)
			return false
		}
	} else {
		log.Infoln("Same Ids")
	}

	return true
}
