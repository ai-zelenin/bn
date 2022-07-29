package parser

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/ai.zelenin/bpmn/pkg/spec"
	"io/ioutil"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFile(fn string) error {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	var defs = new(spec.Definitions)
	err = xml.Unmarshal(data, defs)
	if err != nil {
		return err
	}
	d, err := json.MarshalIndent(defs, "", "\t")
	fmt.Println(string(d))
	return nil
}
