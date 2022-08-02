package spec

import "encoding/xml"

type BPMNDiagram struct {
	XMLName   xml.Name `xml:"BPMNDiagram"`
	Text      string   `xml:",chardata"`
	Id        string   `xml:"id,attr"`
	BPMNPlane struct {
		Text        string `xml:",chardata"`
		Id          string `xml:"id,attr"`
		BpmnElement string `xml:"bpmnElement,attr"`
		BPMNShape   []struct {
			Text            string `xml:",chardata"`
			Id              string `xml:"id,attr"`
			BpmnElement     string `xml:"bpmnElement,attr"`
			IsHorizontal    string `xml:"isHorizontal,attr"`
			IsMarkerVisible string `xml:"isMarkerVisible,attr"`
			IsExpanded      string `xml:"isExpanded,attr"`
			Bounds          struct {
				Text   string `xml:",chardata"`
				X      string `xml:"x,attr"`
				Y      string `xml:"y,attr"`
				Width  string `xml:"width,attr"`
				Height string `xml:"height,attr"`
			} `xml:"Bounds"`
			BPMNLabel struct {
				Text   string `xml:",chardata"`
				Bounds struct {
					Text   string `xml:",chardata"`
					X      string `xml:"x,attr"`
					Y      string `xml:"y,attr"`
					Width  string `xml:"width,attr"`
					Height string `xml:"height,attr"`
				} `xml:"Bounds"`
			} `xml:"BPMNLabel"`
		} `xml:"BPMNShape"`
		BPMNEdge []struct {
			Text        string `xml:",chardata"`
			Id          string `xml:"id,attr"`
			BpmnElement string `xml:"bpmnElement,attr"`
			Waypoint    []struct {
				Text string `xml:",chardata"`
				X    string `xml:"x,attr"`
				Y    string `xml:"y,attr"`
			} `xml:"waypoint"`
			BPMNLabel struct {
				Text   string `xml:",chardata"`
				Bounds struct {
					Text   string `xml:",chardata"`
					X      string `xml:"x,attr"`
					Y      string `xml:"y,attr"`
					Width  string `xml:"width,attr"`
					Height string `xml:"height,attr"`
				} `xml:"Bounds"`
			} `xml:"BPMNLabel"`
		} `xml:"BPMNEdge"`
	} `xml:"BPMNPlane"`
}
