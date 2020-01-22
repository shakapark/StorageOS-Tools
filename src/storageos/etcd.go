package storageos

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	etcdKeyPrefix = "storageos/nameidx/nodes/"
)

type nameIndex struct {
	Prefix string `json:"prefix"`
	Key    string `json:"key"`
	ID     string `json:"objectID"`
}

func (ni *nameIndex) getID() string {
	return ni.ID
}

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

	if len(node.Kvs) == 0 {
		return "", errors.New("Node " + hostname + " don't exist in etcd")
	}

	var tmp nameIndex
	err = json.Unmarshal([]byte(node.Kvs[0].Value), &tmp)
	if err != nil {
		return "", errors.New("Parsing Json: " + error.Error(err))
	}

	return (&tmp).getID(), nil
}
