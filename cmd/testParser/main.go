package main

import (
	"github.com/gradecak/dataflow/pkg/parser"
	log "github.com/sirupsen/logrus"
	"os"
)

var data string = `
id: poo
require:
  - name: ayy
    type: lmao
actions:
  Test:
    type: t
    tag: a
    run: g
constraints:
  - border: true
`

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	df := parser.ParseDataflow([]byte(data))
	log.Debug(df)
}
