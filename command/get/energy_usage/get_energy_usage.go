package getEnergyUsage

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getEnergyUsageInterface interface {
	Run() (whInSession int, whAccumulated int, err error)
	parseResponse(response string) (whInSession int, whAccumulated int, err error)
}

type getEnergyUsage struct {
	getEnergyUsageInterface
	command.Command
}

func (c getEnergyUsage) Run() (whInSession int, whAccumulated int, err error) {
	response, err := c.SendRequest()

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

func New(host string) getEnergyUsage {
	c := new(getEnergyUsage)
	c.Host = host
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
