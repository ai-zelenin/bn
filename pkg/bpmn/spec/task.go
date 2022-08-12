package spec

type Task struct {
	Text                             string                           `xml:",chardata"`
	Id                               string                           `xml:"id,attr"`
	Name                             string                           `xml:"name,attr"`
	Documentation                    string                           `xml:"documentation"`
	Incoming                         []string                         `xml:"incoming"`
	Outgoing                         []string                         `xml:"outgoing"`
	MultiInstanceLoopCharacteristics MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
}

func (t *Task) ID() string {
	return t.Id
}

func (t *Task) Definition() *Task {
	return t
}
