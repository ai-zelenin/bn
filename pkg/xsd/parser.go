package xsd

import (
	"aqwari.net/xml/xsd"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Parser struct {
	cfg Config
}

func NewParser(cfg Config) *Parser {
	return &Parser{
		cfg: cfg,
	}
}

func (p *Parser) ParseFile(glob string) ([]xsd.Schema, error) {
	fns, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}
	var data [][]byte
	for _, fn := range fns {
		fmt.Println(fn)
		d, err := ioutil.ReadFile(fn)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	schemaList, err := xsd.Parse(data...)
	if err != nil {
		return nil, err
	}
	return schemaList, nil
}
