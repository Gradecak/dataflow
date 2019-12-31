package consent

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
)

type ConsentStatus uint32

const (
	REVOKED   ConsentStatus = 0
	GRANTED   ConsentStatus = 1
	PAUSED    ConsentStatus = 2
	UNDEFINED ConsentStatus = 3
)

func (status ConsentStatus) String() string {
	statuses := [...]string{
		"REVOKED",
		"GRANTED",
		"PAUSED",
		"UNDEFINED",
	}
	if status < REVOKED || status > PAUSED {
		return "N/A"
	}
	return statuses[status]
}

func (status ConsentStatus) MarshalBinary() ([]byte, error) {
	if status < REVOKED || status > PAUSED {
		log.Error("Value outside of enum range")
		return nil, errors.New("Value outside of enum range")
	}
	barr := make([]byte, 4)
	binary.LittleEndian.PutUint32(barr, uint32(status))
	return barr, nil
}

func (status *ConsentStatus) UnmarshalBinary(data []byte) error {
	intStatus := binary.LittleEndian.Uint32(data)
	*status = ConsentStatus(intStatus)
	return nil
}

type ConsentMessage struct {
	Id     string
	Status ConsentStatus
}

func (msg ConsentMessage) isValid() bool {
	return msg.Id != "" &&
		(msg.Status != UNDEFINED ||
			!(msg.Status < REVOKED) ||
			!(msg.Status > UNDEFINED))
}

/* **********************************************************************
 * Encoders and Decoders for transmitting Consent messages over the network
 * ********************************************************************* */
func (msg ConsentMessage) WireEncode() []byte {
	str, err := json.Marshal(msg)
	if err != nil {
		log.WithFields(log.Fields{"msg": msg}).Error()
	}
	return []byte(str)
}

func (msg *ConsentMessage) WireDecode(data []byte) {
	log.Debug("Decoding ", string(data))
	err := json.Unmarshal(data, msg)
	if err != nil {
		log.Fatal("Failed unmarshaling data ", err)
	}
}
