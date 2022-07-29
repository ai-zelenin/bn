package spec

type TimerEventDefinition struct {
	Text         string       `xml:",chardata"`
	ID           string       `xml:"id,attr"`
	TimeDuration TimeDuration `xml:"timeDuration"`
}

type TimeDuration struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}
