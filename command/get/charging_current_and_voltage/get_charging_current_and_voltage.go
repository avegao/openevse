package getChargingCurrentAndVoltage

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getChargingCurrentAndVoltageCommandInterface interface {
	Run() (milliAmps int, milliVolts int, err error)
	parseResponse(response string) (milliAmps int, milliVolts int, err error)
}

type getChargingCurrentAndVoltageCommand struct {
	getChargingCurrentAndVoltageCommandInterface
	command.Command
}

func (c getChargingCurrentAndVoltageCommand) Run() (milliAmps int, milliVolts int, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getChargingCurrentAndVoltageCommand) parseResponse(response string) (milliAmps int, milliVolts int, err error) {
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

func New(host string) getChargingCurrentAndVoltageCommand {
	c := new(getChargingCurrentAndVoltageCommand)
	c.Host = host
	c.Type = command.GetChargingCurrentAndVoltage

	return *c
}
