package getOverTemperatureThresholds

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getOverTemperatureThresholdsInterface interface {
	Run() (ambient float32, ir float32, err error)
	parseResponse(response string) (ambient float32, ir float32, err error)
}

type getOverTemperatureThresholds struct {
	getOverTemperatureThresholdsInterface
	command.Command
}

func (c getOverTemperatureThresholds) Run() (ambient float32, ir float32, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getOverTemperatureThresholds) parseResponse(response string) (ambient float32, ir float32, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		var ambientInt, irInt int

		if ambientInt, err = util.ParseInt(split[1]); err != nil {
			return
		}

		if irInt, err = util.ParseInt(split[2]); err != nil {
			return
		}

		ambient = float32(ambientInt) / float32(10)
		ir = float32(irInt) / float32(10)
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New(host string) getOverTemperatureThresholds {
	c := new(getOverTemperatureThresholds)
	c.Host = host
	c.Type = command.GetOverTemperatureThresholds

	return *c
}
