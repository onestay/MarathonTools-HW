package gateway

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/onestay/MarathonTools-HW/serial"
)

var (
	state = 0
)

func HandleButtonCommand(bc <-chan serial.ButtonCommand, apiURL string) {
	for {
		but := <-bc
		if but.Command == "PRESSED" {
			timerCommand(but.Config, apiURL)
		}
	}
}

func timerCommand(conf serial.ButtonConfig, api string) {
	if state == 2 {
		err := startTimer(api)
		if err != nil {
			fmt.Println("Error while starting timer: ", err)
		}
	} else if state == 0 {
		err := stopPlayer(conf.Player, api)
		if err != nil {
			fmt.Println("Error while stopping player: ", err)
		}
	}

}

func startTimer(uri string) error {
	u := url.URL{Scheme: "http", Host: uri, Path: "/timer/start"}

	res, err := http.Post(u.String(), "", nil)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("non 200 status code: %v ", res.StatusCode)
	}

	return nil
}

func stopPlayer(player int, uri string) error {
	u := url.URL{Scheme: "http", Host: uri, Path: "/timer/player/finish/" + string(player)}

	res, err := http.Post(u.String(), "", nil)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("non 200 status code: %v ", res.StatusCode)
	}

	return nil

}
