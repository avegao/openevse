package getFaultCounters

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/util"
)

type getFaultCountersCommandInterface interface {
	Run() (gfdi int, noGround int, stuckRelay int, err error)
	parseResponse(response string) (gfdi int, noGround int, stuckRelay int, err error)
}

type getFaultCountersCommand struct {
	getFaultCountersCommandInterface
	command.Command
}

func (c getFaultCountersCommand) Run() (gfdi int, noGround int, stuckRelay int, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getFaultCountersCommand) parseResponse(response string) (gfdi int, noGround int, stuckRelay int, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		if gfdi, err = util.ParseHexInt(split[1]); err != nil {
			return
		}

		if noGround, err = util.ParseHexInt(split[2]); err != nil {
			return
		}

		if stuckRelay, err = util.ParseHexInt(split[3]); err != nil {
			return
		}
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New(host string) getFaultCountersCommand {
	c := new(getFaultCountersCommand)
	c.Host = host
	c.Type = command.GetFaultCounters

	return *c
}
