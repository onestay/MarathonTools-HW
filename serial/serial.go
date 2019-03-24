package serial

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/tarm/serial"
)

// we are verifiying if we are connected to the right device by sending a ping
func verifyDevice(s *serial.Port, deviceName string) error {
	_, err := s.Write([]byte("PING\n"))
	if err != nil {
		return err
	}

	// read the response
	l, _, err := bufio.NewReader(s).ReadLine()
	if err != nil {
		return err
	}

	d, _ := parseResponse(string(l))
	if d == deviceName {
		return nil
	}

	return fmt.Errorf("wrong device. Got %v, expected %v", d, deviceName)
}

// all responses are structured like <DEVICE_NAME>_<COMMAND>
func parseResponse(response string) (device, command string) {
	parts := strings.Split(response, "_")
	return parts[0], parts[1]
}
