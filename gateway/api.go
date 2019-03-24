package gateway

import (
	"github.com/onestay/MarathonTools-HW/serial"
)

func HandleButtonCommand(bc <-chan serial.ButtonCommand, apiURL string) {
	for {
		but := <-bc
		if but.Command == "PRESSED" {
			timerStop(but.Config, apiURL)
		}
	}
}

func timerStop(conf serial.ButtonConfig, api string) {

}
