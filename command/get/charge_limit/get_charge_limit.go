package getChargeLimit

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getChargeLimitCommandInterface interface {
	Run() (kwh int, err error)
	parseResponse(response string) (kwh int, err error)
}

type getChargeLimitCommand struct {
	getChargeLimitCommandInterface
	command.Command
}

func (c getChargeLimitCommand) Run() (kwh int, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getChargeLimitCommand) parseResponse(response string) (kwh int, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		kwh, err = util.ParseInt(split[1])
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New(host string) getChargeLimitCommand {
	c := new(getChargeLimitCommand)
	c.Host = host
	c.Type = command.GetChargeLimit

	return *c
}
