package currentCapacityRangeInAmps

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getCurrentCapacityRangeInAmpsInterface interface {
	Run(response string) (firmwareVersion string, protocolVersion string, err error)
	parseResponse(response string) (firmwareVersion string, protocolVersion string, err error)
}

type getCurrentCapacityRangeInAmps struct {
	getCurrentCapacityRangeInAmpsInterface
	command.Command
}

func (c getCurrentCapacityRangeInAmps) Run(host string) (minAmps int, maxAmps int, err error) {
	c.Type = command.GetCurrentCapacityRangeInAmps

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getCurrentCapacityRangeInAmps) parseResponse(response string) (minAmps int, maxAmps int, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		if minAmps, err = util.ParseInt(split[1]); err != nil {
			return
		}

		if maxAmps, err = util.ParseInt(split[2]); err != nil {
			return
		}
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New() getCurrentCapacityRangeInAmps {
	c := new(getCurrentCapacityRangeInAmps)
	c.Type = command.GetCurrentCapacityRangeInAmps

	return *c
}
