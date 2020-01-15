package main

import (
	"net/http"
	"strconv"

	"github.com/shakapark/storageos-tools/src/config"

	promlog "github.com/prometheus/common/log"
)

var log promlog.Logger

func init() {
	log = promlog.Base()
}

func checkNodeIDHandler(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig().GetCheckIDConfig()
	success := checkNodeID(cfg.GetEtcdURLs(), cfg.GetHostname(), cfg.GetPathFile())
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	log.Infoln("Starting StorageOS Tools Server")
	cfg := config.GetConfig().GetServerConfig()

	http.HandleFunc("/check/node/id", checkNodeIDHandler)
	log.Fatal(http.ListenAndServe(cfg.GetListenAddress().String()+":"+strconv.Itoa(cfg.GetPort()), nil))
}
