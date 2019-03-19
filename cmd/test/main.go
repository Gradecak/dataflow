package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"pkg/consent"
)

func logInit() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	logInit()
}
