package model

import (
	"encoding/xml"
	"github.com/ai.zelenin/bpmn/pkg/bpmn/spec"
	"io/ioutil"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFile(fn string) (*spec.Definitions, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	var defs = new(spec.Definitions)
	err = xml.Unmarshal(data, defs)
	if err != nil {
		return nil, err
	}
	return defs, nil
}
