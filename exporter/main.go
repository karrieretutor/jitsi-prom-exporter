package main

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	uuid "github.com/satori/go.uuid"
	xmpp "gosrc.io/xmpp"
	stanza "gosrc.io/xmpp/stanza"
)

type iSig int

const (
	iExit iSig = iota
	iFail
)

var osSignals = make(chan os.Signal, 1)
var signals = make(chan iSig)

var jvbbrewery string
var cm *xmpp.StreamManager

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
	xmppUser, ok := os.LookupEnv("XMPP_USER")
	if !ok {
		fmt.Println("No user specified, failing")
		os.Exit(2)
	}

	xmppPw, ok := os.LookupEnv("XMPP_PW")
	if !ok {
		fmt.Println("No liveness password specified, failing")
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

	internalMucDomain, ok := os.LookupEnv("INTERNAL_MUC_DOMAIN")
	if !ok {
		fmt.Println("internal muc domain not specified")
		os.Exit(2)
	}
	jvbbrewery = breweryroom + "@" + internalMucDomain

	jid := xmppUser + "@" + xmppAuthDomain
	fmt.Printf("jid: %s\n", jid)
	address := xmppServer + ":" + xmppPort
	config := xmpp.Config{
		Address:      address,
		Jid:          jid,
		Password:     xmppPw,
		StreamLogger: os.Stdout,
		Insecure:     true,
		TLSConfig:    &tls.Config{InsecureSkipVerify: true},
	}

	router := xmpp.NewRouter()
	router.HandleFunc("message", handleMessage)
	router.HandleFunc("iq", handleIq)
	router.HandleFunc("presence", handlePresence)

	stanza.TypeRegistry.MapExtension(stanza.PKTPresence, xml.Name{Space: "http://jitsi.org/protocol/colibri", Local: "stats"}, Stats{})
	stanza.TypeRegistry.MapExtension(stanza.PKTPresence, xml.Name{Space: "http://jabber.org/protocol/muc#user", Local: "x"}, User{})

	go connectClient(config, router)

	//keep process running
	for s := range signals {
		switch s {
		case iExit:
			shutdown()
			os.Exit(0)
		case iFail:
			shutdown()
			os.Exit(1)
		}
	}
}

func shutdown() {
	cm.Stop()
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

	fmt.Println(presence)

	//check extensions
	for _, e := range presence.Extensions {
		switch extension := e.(type) {
		case *Stats:
			fmt.Printf("\nstats: \n%v\n", extension.Stats)
		case *User:
			for _, i := range extension.Items {
				fmt.Printf("extension Jid: %s\n", i.Jid)
			}
		default:
			fmt.Printf("Found unknown extension: %T\n", extension)
		}
	}
	// err := presence.UnmarshalXML(&xml.Decoder{
	// 	Strict: true,
	// }, xml.StartElement{

	// 	Name: xml.Name{
	// 		Local: "stats",
	// 		Space: "http://jitsi.org/protocol/colibri",
	// 	},
	// })
	// if err != nil {
	// 	fmt.Printf("Encountered error while unmarshalling presence xml: %s", err.Error())
	// }
}

func postConnect(s xmpp.Sender) {
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
	client, err := xmpp.NewClient(c, r)
	if err != nil {
		fmt.Printf("unable to create client: %s\n", err.Error())
		signals <- iFail
	}

	fmt.Println("starting streammanger")

	cm = xmpp.NewStreamManager(client, postConnect)
	err = cm.Run()
	if err != nil {
		fmt.Printf("xmpp connection manager returned with error: %s\n", err.Error())
		signals <- iFail
	}
}
