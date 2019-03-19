package main

import (
	"github.com/gradecak/dataflow/pkg/consent"
	log "github.com/sirupsen/logrus"
	"os"
)

func logInit() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func Test() {
	conf := consent.ConsentStoreConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	store := consent.Init(conf)

	store.SetConsent(consent.ConsentMessage{Id: "Maki", Status: REVOKED})
	store.GetConsent("Maki")
}

func main() {
	logInit()
	Test()
}
