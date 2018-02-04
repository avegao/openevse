package getAuthLockState

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
)

type getAuthLockStateInterface interface {
	Run() (locked bool, err error)
	parseResponse(response string) (locked bool, err error)
}

type getAuthLockStateCommand struct {
	getAuthLockStateInterface
	command.Command
}

func (c getAuthLockStateCommand) Run() (locked bool, err error) {
	return false, errors.New("openevse - this method seem that is not implement in the device")

	//response, err := c.SendRequest()
	//
	//if err != nil {
	//	return
	//}
	//
	//return c.parseResponse(response.Response)
}

func (c getAuthLockStateCommand) parseResponse(response string) (locked bool, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		return
	case command.FailureResponse:
		err = errors.New("invalid request")
	default:
		err = errors.New("invalid request - unknown error")
	}

	return
}

func New(host string) getAuthLockStateCommand {
	c := new(getAuthLockStateCommand)
	c.Host = host
	c.Type = command.GetVersion

	return *c
}
