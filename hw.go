package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/onestay/MarathonTools-HW/serial"

	"github.com/onestay/MarathonTools-HW/gateway"
)

var (
	dispSerial, buttonConfig string
	apiURL, httpPort         string
)

func init() {
	flag.StringVar(&dispSerial, "dispSerial", "", "path to the serial device of the display")
	flag.StringVar(&apiURL, "api", "localhost:3000", "url to MarathonTools API")
	flag.StringVar(&buttonConfig, "buttonConfig", "./buttonconfig.json", "File for the button configuration")
	flag.StringVar(&httpPort, "port", ":3002", "port for the webserver")
	flag.Parse()
	buttonConfig, _ = filepath.Abs(buttonConfig)

}

func main() {
	timeUpdate := make(chan float64)
	butCom := make(chan serial.ButtonCommand)
	gateway.OpenWSConnection(apiURL, timeUpdate)
	err := serial.OpenButtonSerial(buttonConfig, butCom)
	if err != nil {
		log.Fatal(err)
	}
	serial.OpenDisplaySerial(dispSerial, timeUpdate)

	fmt.Println("done")
}
