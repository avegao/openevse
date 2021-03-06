package getEvConnectState

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/command/ev_connect_state"
	"github.com/avegao/openevse/util"
)

type getEvConnectStateCommandInterface interface {
	Run() (state evConnectState.EvConnectState, err error)
	parseResponse(response string) (state evConnectState.EvConnectState, err error)
}

type getEvConnectStateCommand struct {
	getEvConnectStateCommandInterface
	command.Command
}

func (c getEvConnectStateCommand) Run() (state evConnectState.EvConnectState, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getEvConnectStateCommand) parseResponse(response string) (state evConnectState.EvConnectState, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		var stateParsed int
		stateParsed, err = util.ParseInt(split[1])

		state = evConnectState.EvConnectState(stateParsed)
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New(host string) getEvConnectStateCommand {
	c := new(getEvConnectStateCommand)
	c.Host = host
	c.Type = command.GetEvConnectState

	return *c
}
