package main

import (
	"flag"
	"fmt"
	"github.com/amaxlab/go-lib/log"
	"github.com/amaxlab/mac-radar/mikrotik"
)

func main() {
	config := NewConfiguration()
	manager := NewMacAddressManager(mikrotik.NewClient(config.MikrotikConfiguration))

	debugArg := flag.Bool("debug", config.Debug, "is bool")
	portArg := flag.Int("port", config.Port, "is int")

	flag.Parse()

	if *debugArg {
		log.Debug.Enable()
	}

	go manager.Watch()

	log.Info.Println(fmt.Sprintf("Listen: %d port", *portArg))
	server := NewWebServer(config.Port, manager)
	err := server.start()
	if err != nil {
		log.Error.Fatal(err)
	}
}
