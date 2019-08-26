package main

import "encoding/xml"

type Stats struct {
	XMLName xml.Name `xml:"http://jitsi.org/protocol/colibri stats"`
	Stats   []Stat   `xml:"stat"`
}

type Stat struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func (s Stats) Namespace() string {
	return s.XMLName.Space
}
