package ecflow_watchman

import (
	"github.com/perillaroc/ecflow-client-go"
	log "github.com/sirupsen/logrus"
)

func GetEcflowStatus(host string, port string) {
	client := ecflow_client.CreateEcflowClient(host, port)
	defer client.Close()

	ret := client.Sync()
	if ret != 0 {
		log.Fatal("sync has error")
	}

	records := client.StatusRecords()
	log.Info("get nodes: ", len(records))
}