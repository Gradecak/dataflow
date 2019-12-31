package main

import (
	"github.com/gradecak/dataflow/pkg/consent"
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
	"os"
)

func logInit() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
}

func Test() {
	sc, err := stan.Connect("test-cluster", "consent-client")
	if err != nil {
		panic("shite")
	}

	msg := consent.ConsentMessage{Id: "Maki", Status: consent.PAUSED}
	sc.Publish("CONSENT", msg.WireEncode())
}

func main() {
	logInit()
	Test()
}
