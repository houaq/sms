package main

import (
	"log"

	smslib "github.com/houaq/sms/lib"
	"github.com/houaq/sms/modem"
)

func main() {
	cfg, err := NewConfig("config.toml")
	if err != nil {
		log.Fatalf("main: Invalid config: %s", err.Error())
	}

	db, err := smslib.InitDB("db.sqlite")
	defer db.Close()
	if err != nil {
		log.Fatalf("main: Error initializing database: %s", err.Error())
	}

	err = modem.InitModem(cfg.ComPort, cfg.BaudRate)
	if err != nil {
		log.Fatalf("main: error initializing to modem. %s", err)
	}
	err = modem.Reset()
	if err != nil {
		log.Fatalf("main: error reseting modem. %s", err)
	}
	smslib.InitWorker()
	err = InitServer(cfg.ServerHost, cfg.ServerPort)
	if err != nil {
		log.Fatalf("main: Error starting server: %s", err.Error())
	}
}
