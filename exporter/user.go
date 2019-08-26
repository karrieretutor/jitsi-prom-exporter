package main

import "encoding/xml"

// User represents the http://jabber.org/protocol/muc#user extension in a limited fashion
// see https://xmpp.org/extensions/xep-0045.html#schemas-user for details, we parse a subset of 'item'
// attributes only
type User struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/muc#user x"`
	Items   []Item   `xml:"item"`
}

// Item represents an item as described in https://xmpp.org/extensions/xep-0045.html#schemas-user
// Currently there are only three attributes parsed: role, jid, affiliation
type Item struct {
	Role        string `xml:"role,attr"`
	Jid         string `xml:"jid,attr"`
	Affiliation string `xml:"affiliation,attr"`
}
