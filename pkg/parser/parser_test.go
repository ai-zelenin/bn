package parser

import (
	"testing"
)

func TestParser_ParseFile(t *testing.T) {
	p := NewParser()
	err := p.ParseFile("../../res/bpmn/complex.bpmn")
	if err != nil {
		panic(err)
	}
}
