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

import (
	"fmt"
	"time"
)

//simple tiny watchdog for the streammanager, unfortunately there is no such
//mechanic in go-xmpp
type connectWatchdog struct {
	timeout   time.Duration
	connected bool
}

func (w *connectWatchdog) watchConnection(signals chan<- iSig) {
	time.Sleep(w.timeout)
	if !w.connected {
		fmt.Printf("Failed to connect within %s\n", w.timeout.String())
		signals <- iFailedToConnect
	}
}
