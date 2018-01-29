package getTimeLimit

import (
	"github.com/avegao/openevse/command"
	"strings"
	"errors"
	"github.com/avegao/openevse/util"
)

type getTimeLimitCommandInterface interface {
	Run(host string) (limit int, err error)
	parseResponse(response string) (limit int, err error)
}

type getTimeLimitCommand struct {
	getTimeLimitCommandInterface
	command.Command
}

func (c getTimeLimitCommand) Run(host string) (limit int, err error) {
	c.Type = command.GetTimeLimit

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getTimeLimitCommand) parseResponse(response string) (limit int, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		if limit, err = util.ParseInt(split[1]); err != nil {
			return
		}
	case command.FailureResponse:
		err = errors.New("invalid request")
	default:
		err = errors.New("invalid request - unknown error")
	}

	return
}

func New() getTimeLimitCommand {
	c := new(getTimeLimitCommand)
	c.Type = command.GetTimeLimit

	return *c
}
