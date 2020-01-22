# StorageOS-Tools

## Configuration

Liste Environment Variables:

|   ENV (* Mandatory)   |        Default Value       | Description                                                                        |
|:---------------------:|:--------------------------:|------------------------------------------------------------------------------------|
| SERVER_LISTEN_ADDRESS |          "0.0.0.0"         | Listen address for go server                                                       |
|   SERVER_LISTEN_PORT  |           "8080"           | Listen port for golang server                                                      |
|      ETCD_URLS *      |                            | List of etcd server<br>Format: "http://node1:port1,http://node2,port2"             |
|    NODE_HOSTNAME *    |                            | Hostname of the host                                                               |
|    NODE_PATHFILE *    |                            | Path of the id file of StorageOS<br>Located on host to /var/lib/storageos/state/id |
|   STORAGEOS_ENDPOINT  | "storageos.storageos:5705" | StorageOS api endpoint                                                             |
|   STORAGEOS_USERNAME  |         "storageos"        | StorageOS api username                                                             |
|  STORAGEOS_PASSWORD * |                            | StorageOS api password                                                             |
