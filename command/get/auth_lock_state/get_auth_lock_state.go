package getAuthLockState

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
)

type getAuthLockStateInterface interface {
	Run(response string) (firmwareVersion string, protocolVersion string, err error)
	parseResponse(response string) (firmwareVersion string, protocolVersion string, err error)
}

type getAuthLockStateCommand struct {
	getAuthLockStateInterface
	command.Command
}

func (c getAuthLockStateCommand) Run(host string) (locked bool, err error) {
	return false, errors.New("openevse - this method seem that is not implement in the device")

	c.Type = command.GetAuthLockState

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getAuthLockStateCommand) parseResponse(response string) (locked bool, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		println(response)
	case command.FailureResponse:
		err = errors.New("invalid request")
	default:
		err = errors.New("invalid request - unknown error")
	}

	return
}

func New() getAuthLockStateCommand {
	c := new(getAuthLockStateCommand)
	c.Type = command.GetVersion

	return *c
}
