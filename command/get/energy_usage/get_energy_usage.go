package getEnergyUsage

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getDelayTimerInterface interface {
	Run(response string) (firmwareVersion string, protocolVersion string, err error)
	parseResponse(response string) (firmwareVersion string, protocolVersion string, err error)
}

type getEnergyUsage struct {
	getDelayTimerInterface
	command.Command
}

func (c getEnergyUsage) Run(host string) (whInSession int, whAccumulated int, err error) {
	c.Type = command.GetEnergyUsage

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getEnergyUsage) parseResponse(response string) (whInSession int, whAccumulated int, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		whInSession, whAccumulated, err = parseResponseFromSplit(split)

		whInSession /= 3600
		whAccumulated /= 1000
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New() getEnergyUsage {
	c := new(getEnergyUsage)
	c.Type = command.GetEnergyUsage

	return *c
}

func parseResponseFromSplit(split []string) (whInSession int, whAccumulated int, err error) {
	if whInSession, err = util.ParseInt(split[1]); err != nil {
		return
	}

	if whAccumulated, err = util.ParseInt(split[2]); err != nil {
		return
	}

	return
}
