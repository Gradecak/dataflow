package provenance

import (
	"errors"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"time"
)

type InvocationId = string

type Provenance struct {
	Pending *cache.Cache    //key value store InvocationId:_task
	Tasks   map[TaskId]Task //map of all available "tasks"(dataflows) in the system
}

func NewProvenance() Provenance {
	cache := cache.New(10*time.Minute, 15*time.Minute)
	return Provenance{Pending: cache}
}

func (p *Provenance) NewTask(t Task) error {
	_, exists := p.Tasks[t.Id]
	if exists {
		log.WithField("task_id", t.Id).Error("Adding a non unique task")
		return errors.New("Task id not unique")
	} else {
		p.Tasks[t.Id] = t
	}
	return nil
}
