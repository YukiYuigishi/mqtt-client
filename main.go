package main

import (
	"flag"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var msgCh = make(chan mqtt.Message)

func main() {
	var (
		port    = flag.Int("p", 1883, "set mqtt port")
		pubMode = flag.Bool("pub", false, "publisher mode")
		subMode = flag.Bool("sub", false, "subscriber mode")
		topic   = flag.String("t", "", "topic")
		message = flag.String("m", "", "message")
	)

	if *pubMode && *subMode || (*pubMode || *subMode) {
		//TODO
		log.Fatal("mode error: you should set one mode parm")
	}

   opts:=mqtt.NewClientOptions()
   opts.AddBroker("tco:)

}

// / MessageHandler
func MessageHandler(c mqtt.Client, m mqtt.Message) {
	msgCh <- m
}
