package main

import (
	"github.com/gradecak/dataflow/pkg/consent"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func logInit() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	// log.SetReportCaller(true)
}

func Test() {
	conf := consent.ConsentConfig{
		ClusterName: "test-cluster",
		RedisConf: consent.ConsentStoreConfig{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
