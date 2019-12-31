package consent

import (
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
)

func connectionLostHandler(_ stan.Conn, reason error) {
	log.Fatal("Connection to NATS-Streaming lost ", reason)
}

func msgReceivedHandler(m *stan.Msg) {
	cm := ConsentMessage{}
	cm.WireDecode([]byte(m.Data))
	log.Debug("Recieved ", cm)

}
