package getCurrentCapacityRangeInAmps

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getCurrentCapacityRangeInAmpsCommandInterface interface {
	Run() (minAmps int, maxAmps int, err error)
	parseResponse(response string) (minAmps int, maxAmps int, err error)
}

type getCurrentCapacityRangeInAmpsCommand struct {
	getCurrentCapacityRangeInAmpsCommandInterface
	command.Command
}

func (c getCurrentCapacityRangeInAmpsCommand) Run() (minAmps int, maxAmps int, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getCurrentCapacityRangeInAmpsCommand) parseResponse(response string) (minAmps int, maxAmps int, err error) {
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

func New(host string) getCurrentCapacityRangeInAmpsCommand {
	c := new(getCurrentCapacityRangeInAmpsCommand)
	c.Host = host
	c.Type = command.GetCurrentCapacityRangeInAmps

	return *c
}
