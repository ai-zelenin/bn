package spec

type TimerEventDefinition struct {
	Text         string       `xml:",chardata"`
	Id           string       `xml:"id,attr"`
	TimeDuration TimeDuration `xml:"timeDuration"`
}

func (t *TimerEventDefinition) ID() string {
	return t.Id
}

type TimeDuration struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}
