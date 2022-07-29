package xsd

import "testing"

func TestCodeGen_Generate(t *testing.T) {
	cfg := Config{
		TargetSpace:  "http://www.omg.org/spec/BPMN/20100524/MODEL",
		OutDir:       "./spec",
		PackageName:  "spec",
		TemplateGlob: "./templates/*",
		Imports: []string{
			"encoding/xml",
		},
	}
	p := NewParser(cfg)

	schemList, err := p.ParseFile("../../res/xsd/BPMN20.xsd")
	if err != nil {
		panic(err)
	}
	cg, err := NewCodeGen(cfg)
	if err != nil {
		panic(err)
	}
	err = cg.Generate(schemList)
	if err != nil {
		panic(err)
	}

}
