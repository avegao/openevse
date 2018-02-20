package getEnergyUsage

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getEnergyUsageCommandInterface interface {
	Run() (whInSession float32, whAccumulated float32, err error)
	parseResponse(response string) (whInSession float32, whAccumulated float32, err error)
}

type getEnergyUsageCommand struct {
	getEnergyUsageCommandInterface
	command.Command
}

func (c getEnergyUsageCommand) Run() (whInSession float32, whAccumulated float32, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getEnergyUsageCommand) parseResponse(response string) (whInSession float32, whAccumulated float32, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		var whInSessionInt, whAccumulatedInt int
		whInSessionInt, whAccumulatedInt, err = parseResponseFromSplit(split)

		whInSession = float32(whInSessionInt)/3600
		whAccumulated = float32(whAccumulatedInt)/1000
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New(host string) getEnergyUsageCommand {
	c := new(getEnergyUsageCommand)
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
