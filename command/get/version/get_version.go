package getVersion

import (
	"github.com/avegao/openevse/command"
	"strings"
	"errors"
)

type getVersionCommandInterface interface {
	Run(response string) (firmwareVersion string, protocolVersion string, err error)
	parseResponse(response string) (firmwareVersion string, protocolVersion string, err error)
}

type getVersionCommand struct {
	getVersionCommandInterface
	command.Command
}

func (c getVersionCommand) Run(host string) (firmwareVersion string, protocolVersion string, err error) {
	c.Type = command.GetVersion

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getVersionCommand) parseResponse(response string) (firmwareVersion string, protocolVersion string, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		firmwareVersion, protocolVersion = parseVersionFromResponseSplit(split)
	case command.FailureResponse:
		err = errors.New("invalid request")
	default:
		err = errors.New("invalid request - unknown error")
	}

	return
}

func New() getVersionCommand {
	c := new(getVersionCommand)
	c.Type = command.GetVersion

	return *c
}

func parseVersionFromResponseSplit(split []string) (firmwareVersion string, protocolVersion string) {
	return split[1], split[2]
}
