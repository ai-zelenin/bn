package spec

type Error struct {
	Text      string `xml:",chardata"`
	Id        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	ErrorCode string `xml:"errorCode,attr"`
}

func (t *Error) ID() string {
	return t.Id
}

type ErrorEventDefinition struct {
	Text     string `xml:",chardata"`
	Id       string `xml:"id,attr"`
	ErrorRef string `xml:"errorRef,attr"`
}

func (t *ErrorEventDefinition) ID() string {
	return t.Id
}
