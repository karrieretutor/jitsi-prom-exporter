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
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	xmpp "github.com/FluuxIO/go-xmpp"
	uuid "github.com/satori/go.uuid"
	stanza "gosrc.io/xmpp/stanza"
)

type iSig int

const (
	iExit iSig = iota
	iFail
	iFailedToConnect
)

var (
	osSignals = make(chan os.Signal, 1)
	signals   = make(chan iSig)
	watchdog  = connectWatchdog{
		timeout:   5 * time.Second,
		connected: false,
	}

	serverChecker *xmpp.ServerCheck

	jvbbrewery string
	cm         *xmpp.StreamManager

	jvbCollector = NewJvbCollector(os.Getenv("JVB_METRIC_NAMESPACE"), os.Getenv("JVB_METRIC_SUBSYSTEM"), 30*time.Second)
)

func init() {
	//xmpp stuff
	stanza.TypeRegistry.MapExtension(stanza.PKTPresence, xml.Name{Space: "http://jitsi.org/protocol/colibri", Local: "stats"}, Stats{})
	stanza.TypeRegistry.MapExtension(stanza.PKTPresence, xml.Name{Space: "http://jabber.org/protocol/muc#user", Local: "x"}, User{})

	prometheus.MustRegister(jvbCollector)
}

func main() {
	//set up os signal listener
	signal.Notify(osSignals, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for s := range osSignals {
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				fmt.Println("Got told by os signal to exit")
				signals <- iExit
			}
		}
	}()

	//read credentials and stuff from environment
	xmppUser, ok := os.LookupEnv("PROMEXP_AUTH_USER")
	if !ok {
		fmt.Println("No user specified, failing")
		os.Exit(2)
	}

	xmppPw, ok := os.LookupEnv("PROMEXP_AUTH_PASSWORD")
	if !ok {
		fmt.Println("No password specified, failing")
		os.Exit(2)
	}

	xmppAuthDomain, ok := os.LookupEnv("XMPP_AUTH_DOMAIN")
	if !ok {
		fmt.Println("no xmpp auth domain specified")
		os.Exit(2)
	}

	xmppPort, ok := os.LookupEnv("XMPP_PORT")
	if !ok || xmppPort == "" {
		xmppPort = "5222"
	}

	xmppServer, ok := os.LookupEnv("XMPP_SERVER")
	if !ok {
		fmt.Println("no xmpp server specified")
		os.Exit(2)
	}

	breweryroom, ok := os.LookupEnv("JVB_BREWERY_MUC")
	if !ok || breweryroom == "" {
		breweryroom = "jvbbrewery"
	}

	internalMucDomain, ok := os.LookupEnv("XMPP_INTERNAL_MUC_DOMAIN")
	if !ok {
		fmt.Println("internal muc domain not specified")
		os.Exit(2)
	}

	//we need a serverchecker
	go func() {
		address := xmppServer + ":" + xmppPort
		for true {
			conn, err := net.DialTimeout("tcp", address, 5*time.Second)
			if err != nil {
				fmt.Printf("Could not connect to server %s: %s\nexiting\n", address, err.Error())
				signals <- iExit
				return
			}

			if tcpconn, ok := conn.(*net.TCPConn); ok {
				_ = tcpconn.Close()
			}
			time.Sleep(10 * time.Second)
		}
	}()

	jvbbrewery = breweryroom + "@" + internalMucDomain

	jid := xmppUser + "@" + xmppAuthDomain
	address := xmppServer + ":" + xmppPort
	config := xmpp.Config{
		TransportConfiguration: xmpp.TransportConfiguration{
			Address:      address,
			Domain:       xmppAuthDomain,
			TLSConfig:    &tls.Config{InsecureSkipVerify: true},
		},
		Jid:          jid,
		Credential:   xmpp.Password(xmppPw),
		StreamLogger: os.Stdout,
		Insecure:     true,
	}

	router := xmpp.NewRouter()
	router.HandleFunc("message", handleMessage)
	router.HandleFunc("iq", handleIq)
	router.HandleFunc("presence", handlePresence)

	go connectClient(config, router)

	//start the watchdog
	go watchdog.watchConnection(signals)

	//start serving prom metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("Unable to serve prom metrics: %s\n", err.Error())
			signals <- iFail
		}
	}()

	//keep process running
	for s := range signals {
		switch s {
		case iExit:
			shutdown()
			os.Exit(0)
		case iFail:
			shutdown()
			os.Exit(1)
		case iFailedToConnect:
			os.Exit(2)
		}
	}
}

func shutdown() {
	//we get a segfault if connection actually never happened
	if cm != nil {
		cm.Stop()
	}
}

func handleMessage(s xmpp.Sender, p stanza.Packet) {

}

func handleIq(s xmpp.Sender, p stanza.Packet) {

}

func handlePresence(s xmpp.Sender, p stanza.Packet) {
	presence, ok := p.(stanza.Presence)
	if !ok {
		fmt.Println("received presence which is no presence, skipping")
		return
	}

	if presence.Get(&Stats{}) && presence.Get(&User{}) {
		var jvbJid string
		var stats *Stats

		//check extensions
		for _, e := range presence.Extensions {
			switch extension := e.(type) {
			case *Stats:
				stats = extension
			case *User:
				for _, i := range extension.Items {
					jvbJid = i.Jid
				}
			}
		}

		//we want to keep track of jvbs across their reconnects of autoscaled sets
		jvbCollector.Update(strings.Split(jvbJid, "@")[0], stats)
	}
}

func postConnect(s xmpp.Sender) {
	watchdog.connected = true

	client, ok := s.(*xmpp.Client)
	if !ok {
		fmt.Println("post connect sender not a client, cannot proceed")
		signals <- iFail
	}

	uuid, _ := uuid.NewV4()
	id := uuid.String()

	//join jvbbrewery room
	presence := stanza.NewPresence(stanza.Attrs{
		To:   jvbbrewery + "/prom-exporter",
		From: client.Session.BindJid,
		Id:   id,
	})

	presence.Extensions = append(presence.Extensions, stanza.MucPresence{})

	err := client.Send(presence)
	if err != nil {
		//err handling
	}
}

func connectClient(c xmpp.Config, r *xmpp.Router) {
	client, err := xmpp.NewClient(&c, r, errorHandler)
	if err != nil {
		fmt.Printf("unable to create client: %s\n", err.Error())
	}

	fmt.Println("starting streammanger")

	cm = xmpp.NewStreamManager(client, postConnect)
	err = cm.Run()
	if err != nil {
		fmt.Printf("xmpp connection manager returned with error: %s\n", err.Error())
		signals <- iFail
		return
	}

	//connection closed we are done
	fmt.Println("XMPP connection closed, exiting.")
	signals <- iExit
	return
}

// If an error occurs, this is used to kill the client
func errorHandler(err error) {
	signals <- iFail
}
