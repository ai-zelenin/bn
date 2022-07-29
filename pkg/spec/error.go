package spec

type Error struct {
	Text      string `xml:",chardata"`
	ID        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	ErrorCode string `xml:"errorCode,attr"`
}

type ErrorEventDefinition struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"id,attr"`
	ErrorRef string `xml:"errorRef,attr"`
}
