package serial

import (
	"fmt"
	"time"

	"github.com/tarm/serial"
)

// OpenDisplaySerial opens the serial connection with the display
func OpenDisplaySerial(dev string, timeUpdate <-chan float64) error {
	c := &serial.Config{Name: dev, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return err
	}

	rTimeUpdate := make(chan float64)

	err = verifyDevice(s, "DISPLAY1")
	if err != nil {
		return err
	}
	go reduceTimerFrequency(rTimeUpdate, timeUpdate)

	sendTime(rTimeUpdate, s)
	return nil
}

func sendTime(time <-chan float64, s *serial.Port) {
	for {
		<-time
		s.Write([]byte(formatDuration(<-time)))
	}
}

func formatDuration(t float64) string {
	d, _ := time.ParseDuration(fmt.Sprintf("%fs", t))
	d = d.Round(time.Second)

	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	return fmt.Sprintf("%02d:%02d:%02d\n", h, m, s)
}

func reduceTimerFrequency(rTimeUpdate chan<- float64, timeUpdate <-chan float64) {
	var latest float64
	go func() {
		for {
			latest = <-timeUpdate
		}
	}()

	for {
		rTimeUpdate <- latest
		time.Sleep(300 * time.Millisecond)
	}

}
