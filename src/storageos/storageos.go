package storageos

import (
	"context"
	"errors"
	"fmt"

	storageos "github.com/storageos/go-api"
	"github.com/storageos/go-api/types"
)

func newStorageOSClient(nodes, username, password string) (*storageos.Client, error) {
	cli, err := storageos.NewVersionedClient(nodes, "1")
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
		return errors.New("Error creating StorageOS cli: " + err.Error())
	}
	ops := types.DeleteOptions{
		Name:      oldID,
		Namespace: "",
		Force:     true,
		Context:   context.Background(),
	}
	err = client.NodeDelete(ops)
	if err != nil {
		return errors.New("Error deleting node: " + err.Error())
	}

	return nil
}

// ListeStorageOSNode Delete Old Node ID in StorageOS API
func ListeStorageOSNode(oldID, nodes, username, password string) error {
	client, err := newStorageOSClient(nodes, username, password)
	if err != nil {
		return errors.New("Error creating StorageOS cli: " + err.Error())
	}

	ops := types.ListOptions{
		// FieldSelector: "",
		// LabelSelector: "",
		// Namespace: "storageos",
	}
	listNodes, err := client.NodeList(ops)
	if err != nil {
		return errors.New("Error list StorageOS nodes: " + err.Error())
	}

	fmt.Println(listNodes)
	return nil
}
