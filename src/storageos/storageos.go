package storageos

import (
	storageos "github.com/storageos/go-api"
	"github.com/storageos/go-api/types"
)

func newStorageOSClient(nodes, username, password string) (*storageos.Client, error) {
	cli, err := storageos.NewClient(nodes)
	if err != nil {
		return nil, err
	}
	cli.SetAuth(username, password)
	return cli, nil
}

// DeleteStorageOSNode Delete Old Node ID in StorageOS API
func DeleteStorageOSNode(oldID, nodes, username, password string) error {
	client, err := newStorageOSClient(nodes, username, password)
	if err != nil {
		return err
	}
	ops := types.DeleteOptions{
		ID: oldID,
	}
	err = client.NodeDelete(ops)
	return err
}
