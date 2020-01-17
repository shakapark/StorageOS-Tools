package storageos

import (
	storageos "github.com/storageos/go-api"
)

func newStorageOSClient(nodes, username, password string) *storageos.Client {
	cli := storageos.NewClient(nodes)
	cli.SetAuth(username, secret)
	return cli
}

// DeleteStorageOSNode Delete Old Node ID in StorageOS API
func DeleteStorageOSNode(oldID, nodes, username, password string) error {
	client := newStorageOSClient(nodes, username, password)
	ops := storageos.DeleteOptions{
		ID: oldID,
	}
	err := client.NodeDelete(ops)
	return err
}
