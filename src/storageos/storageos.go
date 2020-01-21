package storageos

import (
	"context"
	"errors"
	"fmt"

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
		return errors.New("Error creating StorageOS cli: " + err.Error())
	}
	ops := types.DeleteOptions{
		Name: oldID,
	}
	err = client.NodeDelete(ops)
	return errors.New("Error deleting node: " + err.Error())
}

// ListeStorageOSNode Delete Old Node ID in StorageOS API
func ListeStorageOSNode(oldID, nodes, username, password string) error {
	client, err := newStorageOSClient(nodes, username, password)
	if err != nil {
		return errors.New("Error creating StorageOS cli: " + err.Error())
	}
	ops := types.ListOptions{
		FieldSelector: "",
		LabelSelector: "",
		Context:       context.Background(),
	}
	listNodes, err := client.NodeList(ops)
	if err != nil {
		return errors.New("Error list StorageOS nodes: " + err.Error())
	}

	fmt.Println(listNodes)
	return nil
}
