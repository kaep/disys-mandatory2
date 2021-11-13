package main

import (
	"log"
	"os"

	"github.com/hashicorp/serf/serf"
	"github.com/pkg/errors"
)

func main() {
	var advertise = "172.17.0.2"
	cluster, err := setupCluster(
		advertise,
		os.Getenv("CLUSTER_ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer cluster.Leave()

}

func setupCluster(advertiseAddr string, clusterAddr string) (*serf.Serf, error) {
	conf := serf.DefaultConfig()
	conf.Init()
	conf.MemberlistConfig.AdvertiseAddr = advertiseAddr

	cluster, err := serf.Create(conf)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create cluster")
	}

	_, err = cluster.Join([]string{clusterAddr}, true)
	if err != nil {
		log.Printf("Couldn't join cluster, starting own: %v\n", err)
	}

	return cluster, nil
}
