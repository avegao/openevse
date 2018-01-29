package getVoltmeterSettings

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getVoltmeterSettingsInterface interface {
	Run(host string) (calefactor int, offset int, err error)
	parseResponse(response string) (limit int, err error)
}

type getVoltmeterSettings struct {
	getVoltmeterSettingsInterface
	command.Command
}

func (c getVoltmeterSettings) Run(host string) (calefactor int, offset int, err error) {
	c.Type = command.GetVoltmeterSettings

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getVoltmeterSettings) parseResponse(response string) (calefactor int, offset int, err error) {
	split := strings.Split(response, " ")

	println(response)

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

func New() getVoltmeterSettings {
	c := new(getVoltmeterSettings)
	c.Type = command.GetVoltmeterSettings

	return *c
}
