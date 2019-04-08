package gateway

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type wsTimeUpdate struct {
	DataType string  `json:"dataType"`
	T        float64 `json:"t"`
	State    int     `json:"state"`
}

// OpenWSConnection opens a connection to the MarathonTools-API websocket
func OpenWSConnection(apiURL string, timeUpdate chan<- float64) error {
	u := url.URL{Scheme: "ws", Host: apiURL, Path: "/ws"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	log.Println("Connected to websocket on ", u.String())

	go readWSMessages(timeUpdate, c)

	return nil
}

func readWSMessages(tu chan<- float64, c *websocket.Conn) {
	for {
		res := wsTimeUpdate{}
		c.ReadJSON(&res)
		if res.DataType == "timeUpdate" {
			tu <- res.T
		} else if res.DataType == "stateUpdate" {
			state = res.State
		}
	}

}
