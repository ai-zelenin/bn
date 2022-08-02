package model

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestBPMNSuite(t *testing.T) {
	suite.Run(t, new(BPMNSuite))
}

type BPMNSuite struct {
	suite.Suite
	p     *Parser
	cases []string
}

func (s *BPMNSuite) SetupSuite() {
	s.p = NewParser()
	s.cases = []string{"../../res/test-cases/lvl1.bpmn"}
}

func (s *BPMNSuite) TestParseAndRun() {
	for _, testCase := range s.cases {
		sp, err := s.p.ParseFile(testCase)
		s.Require().Nil(err)
		s.Require().NotNil(sp)
		bpmn, err := NewBPMN(sp)
		s.Require().Nil(err)
		s.Require().NotNil(bpmn)
		for _, process := range bpmn.Processes {
			ctx := context.Background()
			err = process.Run(ctx, "", func(ctx context.Context, el Element) error {
				fmt.Printf("%s -> %s %T \n", process.ID, el.ID(), el)
				return nil
			})
			s.Require().Nil(err)
		}
	}
}
