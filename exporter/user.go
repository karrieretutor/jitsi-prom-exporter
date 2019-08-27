/* 
 *  Copyright 2019 karriere tutor GmbH
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  
 *  	http://www.apache.org/licenses/LICENSE-2.0
 *  
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *  
 */

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
