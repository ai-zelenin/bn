package model

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
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
	s.cases = []string{"../../../res/bpmn/lvl1.bpmn"}
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
			handler := NewCompositeElementHandler()
			handler.SetHandleError(func(ctx context.Context, elID string, err error) (isContinue bool) {
				fmt.Printf("%s %v", elID, err)
				return true
			})
			handler.AddGlobalHandleElementFunc(func(ctx context.Context, el Element) error {
				fmt.Printf(" %s %T \n", el.ID(), el)
				time.Sleep(time.Millisecond * 500)
				if el.ID() == "Activity_4" {
					time.Sleep(time.Millisecond * 3000)
				}
				return nil
			})
			exec, err := bpmn.LoadExecutableModel(ctx, process.ID())
			s.Require().Nil(err)
			exec.Exec(ctx, []string{process.StartEvent.ID()}, handler)
			time.Sleep(time.Millisecond * 500)
		}
		time.Sleep(time.Second * 10)
	}
}
