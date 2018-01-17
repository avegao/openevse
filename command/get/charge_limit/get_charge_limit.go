package getChargeLimit

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getChargeLimitInterface interface {
	Run(response string) (firmwareVersion string, protocolVersion string, err error)
	parseResponse(response string) (firmwareVersion string, protocolVersion string, err error)
}

type getChargeLimit struct {
	getChargeLimitInterface
	command.Command
}

func (c getChargeLimit) Run(host string) (kwh int, err error) {
	c.Type = command.GetChargeLimit

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getChargeLimit) parseResponse(response string) (kwh int, err error) {
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

func New() getChargeLimit {
	c := new(getChargeLimit)
	c.Type = command.GetChargeLimit

	return *c
}
