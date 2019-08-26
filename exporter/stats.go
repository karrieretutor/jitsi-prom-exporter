package main

import "encoding/xml"

// Stats Container for individual stats as they are broadcasted by the JVBs.
// Stats is then registered as gosrc.io/xmpp/stanza.PresExetension which unmarshals
// the xml coming in via xmpp into these Stat structs
type Stats struct {
	XMLName xml.Name `xml:"http://jitsi.org/protocol/colibri stats"`
	Stats   []Stat   `xml:"stat"`
}

// Stat Individual stat containing a single metric, such as Name="bit_rate_upload", Value="0"
type Stat struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
