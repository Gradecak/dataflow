package provenance

import (
	"fmt"
	// log "github.com/sirupsen/logrus"
)

type Triple struct {
	Src  string
	Edge string
	Dest string
}

func (t Triple) String() string {
	return fmt.Sprintf("(%s,%s,%s)", t.Src, t.Edge, t.Dest)
}

// vertex is encoded as triple
type Vertex interface {
	ToTriple() Triple
}

// a graph is just a set of triples
type Graph interface {
	ToGraph() []Triple
}

type TaskId = string

// full Task structure for generating the provenance graph
type Task struct {
	Id          TaskId
	Predecessor *Task
	Subtasks    []Subtask
}

type Subtask struct {
	Type OpType
	Tag  string
}

func (t Task) ToGraph() []Triple {
	trips := []Triple{}

	//predecessor to triple
	if t.Predecessor != nil {
		trips = append(trips, Triple{t.Id, "predecessor", t.Predecessor.Id})
	}

	//subtasks to triples
	for _, st := range t.Subtasks {
		trips = append(trips, Triple{t.Id, st.Type.String(), st.Tag})
	}

	return trips
}

type OpType uint8

const (
	READ      OpType = 0
	WRITE     OpType = 1
	OPERATION OpType = 2
)

func (t OpType) String() string {
	types := [...]string{
		"READ",
		"WRITE",
		"OPERATION",
	}
	if t < READ || t > OPERATION {
		return "NA"
	}
	return types[t]
}

/* ********************************************************************** */
