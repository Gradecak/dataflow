package consent

import (
	"errors"
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

func StoreInit(conf ConsentStoreConfig) Store {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Error("Error Connecting to Redis instance: ", err.Error())
	}
	return Store{client: client}
}

func (store Store) setConsent(msg ConsentMessage) {
	valid := msg.isValid()
	if valid {
		log.WithFields(log.Fields{"id": msg.Id, "status": msg.Status}).Debug()

		err := store.client.Set(msg.Id, msg.Status, 0).Err()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (store Store) getConsent(id string) (ConsentStatus, error) {
	log.WithFields(log.Fields{"id": id}).Debug("Fetching From Store..")
	val, err := store.client.Get(id).Result()

	//if the value is not in redis store
	if err == redis.Nil {
		log.WithFields(log.Fields{"id": id}).Info("key does not exist")
		return GRANTED, nil
	}
	if err != nil {
		log.WithFields(log.Fields{"id": id}).Error(err.Error())
		return UNDEFINED, err
	}

	status := UNDEFINED
	err = status.UnmarshalBinary([]byte(val))
	// status, err := strconv.Atoi(val)
	if err != nil {
		log.WithFields(log.Fields{"status": val}).Error("Fetched status not in valid ConsentStatus range")
		return UNDEFINED, errors.New("Cannot unmarshal Stored Value")
	}
	log.WithFields(log.Fields{"status": status}).Debug("Fetched Status:")
	return ConsentStatus(status), nil
}
