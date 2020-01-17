package storageos

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	etcdKeyPrefix = "storageos/nameidx/nodes/"
)

// NewClientETCD Create ETCD Client Object
func NewClientETCD(baseURLs []string, username, password string) (*clientv3.Client, error) {
	cfg := clientv3.Config{
		Endpoints:   baseURLs,
		DialTimeout: 5 * time.Second,
		// TLS: ,
		// Username:  username,
		// Password:  password,
	}

	c, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GetETCDNodeID Get NodeID in ETCD with Hostname
func GetETCDNodeID(c *clientv3.Client, hostname string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kv := clientv3.NewKV(c)
	node, err := kv.Get(ctx, etcdKeyPrefix+hostname)
	if err != nil {
		return "", errors.New("Can't get ETCD key: " + error.Error(err))
	}

	var tmp nameIndex
	err = json.Unmarshal([]byte(node.Kvs[0].Value), &tmp)
	if err != nil {
		return "", errors.New("Parsing Json: " + error.Error(err))
	}

	return (&tmp).getID(), nil
}

// GetFileID Get NodeID in local file
func GetFileID(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.New("Opening File: " + error.Error(err))
	}
	defer f.Close()

	b1 := make([]byte, 64)
	n1, err := f.Read(b1)
	if err != nil {
		return "", errors.New("Reading File: " + error.Error(err))
	}

	if b1[n1-1] == 10 {
		b1 = b1[:(n1 - 1)]
	}

	return string(b1), nil
}

// ReplaceFileID Replace ID in file
func ReplaceFileID(path, newID string) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.New("Opening File: " + error.Error(err))
	}
	defer f.Close()

	_, err = f.WriteString(newID + "\n")
	if err != nil {
		return errors.New("Writing File: " + error.Error(err))
	}

	return nil
}
