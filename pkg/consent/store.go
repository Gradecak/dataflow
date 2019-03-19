package consent

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type ConsentStoreConfig struct {
	Addr     string
	Password string
	DB       int
}

type Store struct {
	client *redis.Client
}

func Init(conf ConsentStoreConfig) Store {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Error(err.Error())
	}
	return Store{client: client}
}

func (store Store) SetConsent(msg ConsentMessage) {
	log.WithFields(log.Fields{"id": msg.id, "status": msg.status}).Debug()

	err := store.client.Set(msg.id, msg.status, 0).Err()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (store Store) GetConsent(id string) ConsentStatus {
	log.WithFields(log.Fields{"id": id}).Debug()
	val, err := store.client.Get(id).Result()

	if err == redis.Nil {
		log.WithFields(log.Fields{"id": id}).Fatal("key does not exist")
	}
	if err != nil {
		log.WithFields(log.Fields{"id": id}).Fatal(err.Error())
	}

	log.WithFields(log.Fields{"val": val}).Debug()
	return ConsentStatus(0)
}
