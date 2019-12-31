package consent

import (
	stan "github.com/nats-io/go-nats-streaming"
	log "github.com/sirupsen/logrus"
)

const SUB_PREFIX = "CONSENT"

type ConsentService struct {
	store      Store
	natsClient stan.Conn
}

type ConsentConfig struct {
	ClusterName string
	RedisConf   ConsentStoreConfig
}

func (cs ConsentService) RunListeners() {

	hndlr := func(m *stan.Msg) {
		cm := ConsentMessage{}
		cm.WireDecode([]byte(m.Data))
		log.Debug("Recieved ", cm)

		//set recievd value in store
		cs.store.setConsent(cm)
	}

	cs.natsClient.Subscribe(SUB_PREFIX, hndlr,
		stan.DurableName("consent-service-durable"))
}

func (cs ConsentService) GetConsent(id string) (ConsentStatus, error) {
	return cs.store.getConsent(id)
}

func Init(conf ConsentConfig) ConsentService {
	cs := ConsentService{}
	cs.store = StoreInit(conf.RedisConf)

	//initialise connection to NATS-Streaming
	sc, err := stan.Connect(conf.ClusterName, "consent-service",
		stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(connectionLostHandler))

	if err != nil {
		log.Error(err.Error())
		panic("Cannot connect to NATS Streaming")
	}

	cs.natsClient = sc
	return cs
}
