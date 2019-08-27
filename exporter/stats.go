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
