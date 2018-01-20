package getChargingCurrentAndVoltage

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getChargingCurrentAndVoltageInterface interface {
	Run(response string) (firmwareVersion string, protocolVersion string, err error)
	parseResponse(response string) (firmwareVersion string, protocolVersion string, err error)
}

type getChargingCurrentAndVoltage struct {
	getChargingCurrentAndVoltageInterface
	command.Command
}

func (c getChargingCurrentAndVoltage) Run(host string) (milliAmps int, milliVolts int, err error) {
	c.Type = command.GetChargingCurrentAndVoltage

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getChargingCurrentAndVoltage) parseResponse(response string) (milliAmps int, milliVolts int, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		if milliAmps, err = util.ParseInt(split[1]); err != nil {
			return
		}

		if milliVolts, err = util.ParseInt(split[2]); err != nil {
			return
		}
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New() getChargingCurrentAndVoltage {
	c := new(getChargingCurrentAndVoltage)
	c.Type = command.GetChargingCurrentAndVoltage

	return *c
}
