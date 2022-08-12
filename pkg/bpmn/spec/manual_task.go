package spec

import "encoding/xml"

type ManualTask struct {
	XMLName xml.Name `xml:"manualTask"`
	Task
}
