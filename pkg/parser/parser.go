package parser

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type DataflowSpec struct {
	Id          string                `yaml:"id"`
	Required    []RequiredParams      `yaml:"require,omitempty"`
	Actions     map[string]Actions    `yaml:"actions"`
	Constraints []map[Constraint]bool `yaml:"constraints"`
}

type RequiredParams struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

type Actions struct {
	Type string `yaml:"type"`
	Tag  string `yaml:"tag"`
	Run  string `yaml:"run"`
}

type Constraint string

const (
	BORDER  Constraint = "border"
	CONSENT Constraint = "consent_check"
)

func ParseDataflow(data []byte) DataflowSpec {
	df := DataflowSpec{}
	err := yaml.Unmarshal(data, &df)

	if err != nil {
		log.Error("cannot unmarshal workflow specification ", err.Error())
	}

	return df
}
