package main

import "encoding/xml"

type Stats struct {
	XMLName  xml.Name `xml:"http://jitsi.org/protocol/colibri stats"`
	InnerXML string   `xml:",innerxml"`
}

func (s Stats) Namespace() string {
	return s.XMLName.Space
}
