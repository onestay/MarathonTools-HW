package serial

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/tarm/serial"
)

type ButtonConfig struct {
	DeviceName string `json:"deviceName"`
	Color      string `json:"color"`
	Player     int    `json:"player"`
	Device     string `json:"device"`
}

type ButtonCommand struct {
	Device  string
	Command string
	Config  ButtonConfig
}

func OpenButtonSerial(configPath string, bc chan<- ButtonCommand) error {
	conf, err := parseConfig(configPath)
	if err != nil {
		return err
	}

	butConf := *conf
	ports := make([]*serial.Port, len(butConf))
	for i, config := range butConf {
		c := &serial.Config{Name: config.Device, Baud: 9600}
		ports[i], err = serial.OpenPort(c)
		if err != nil {
			return err
		}

		// err := verifyDevice(ports[i], config.DeviceName)
		// if err != nil {
		// 	return err
		// }
	}

	readButtonInputs(ports, bc, butConf)
	return nil
}

func readButtonInputs(ports []*serial.Port, bc chan<- ButtonCommand, buttonConfigs []ButtonConfig) error {
	for _, port := range ports {
		go func(p *serial.Port) {
			scanner := bufio.NewScanner(p)
			for scanner.Scan() {
				d, c := parseResponse(scanner.Text())

				var buttonConfig ButtonConfig
				for _, config := range buttonConfigs {
					if config.DeviceName == d {
						buttonConfig = config
					}
				}

				command := ButtonCommand{Device: d, Command: c, Config: buttonConfig}

				bc <- command
			}
		}(port)
	}

	return nil
}

func parseConfig(cp string) (*[]ButtonConfig, error) {
	f, err := os.Open(cp)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var bc []ButtonConfig

	err = json.NewDecoder(f).Decode(&bc)
	if err != nil {
		return nil, err
	}

	return &bc, nil
}
