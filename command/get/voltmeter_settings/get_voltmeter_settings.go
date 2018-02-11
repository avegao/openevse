package getVoltmeterSettings

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getVoltmeterSettingsInterface interface {
	Run() (calefactor int, offset int, err error)
	parseResponse(response string) (calefactor int, offset int, err error)
}

type getVoltmeterSettings struct {
	getVoltmeterSettingsInterface
	command.Command
}

func (c getVoltmeterSettings) Run() (calefactor int, offset int, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getVoltmeterSettings) parseResponse(response string) (calefactor int, offset int, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		if calefactor, err = util.ParseInt(split[1]); err != nil {
			return
		}

		if offset, err = util.ParseInt(split[1]); err != nil {
			return
		}
	case command.FailureResponse:
		err = errors.New("invalid request")
	default:
		err = errors.New("invalid request - unknown error")
	}

	return
}

func New(host string) getVoltmeterSettings {
	c := new(getVoltmeterSettings)
	c.Host = host
	c.Type = command.GetVoltmeterSettings

	return *c
}
