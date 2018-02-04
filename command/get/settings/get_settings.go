package getSettings

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getSettingsCommandInterface interface {
	Run() (amperes int, flags interface{}, err error)
	parseResponse(response string) (amperes int, flags int, err error)
}

type getSettingsCommand struct {
	getSettingsCommandInterface
	command.Command
}

func (c getSettingsCommand) Run() (amperes int, flags interface{}, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getSettingsCommand) parseResponse(response string) (amperes int, flags int, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		if amperes, err = util.ParseInt(split[1]); err != nil {
			return
		}

		if flags, err = util.ParseHexInt(split[2]); err != nil {
			return
		}
	case command.FailureResponse:
		err = errors.New("invalid request")
	default:
		err = errors.New("invalid request - unknown error")
	}

	return
}

func New(host string) getSettingsCommand {
	c := new(getSettingsCommand)
	c.Host = host
	c.Type = command.GetSettings

	return *c
}
