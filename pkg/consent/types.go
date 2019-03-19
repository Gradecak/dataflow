package consent

type ConsentStatus int

const (
	REVOKED ConsentStatus = 0
	GRANTED ConsentStatus = 1
	PAUSED  ConsentStatus = 2
)

func (status ConsentStatus) String() string {
	statuses := [...]string{
		"REVOKED",
		"GRANTED",
		"PAUSED",
	}
	if status < REVOKED || status > PAUSED {
		return "N/A"
	}
	return statuses[status]
}

type ConsentMessage struct {
	Id     string
	Status ConsentStatus
}
