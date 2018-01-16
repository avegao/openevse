package getAmmeterSettings

import (
	"github.com/avegao/openevse/command"
	"strings"
	"github.com/avegao/openevse/util"
)

type getAmmeterSettingsCommandInterface interface {
	Run(response string) (string, error)
	parseResponse(response string) (error)
}

type getAmmeterSettingsCommand struct {
	getAmmeterSettingsCommandInterface
	command.Command
}

func (c getAmmeterSettingsCommand) Run(host string) (currentScaleFactor int, currentOffset int, err error) {
	c.Type = command.GetAmmeterSettings

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return parseResponse(response.Response)
}

func New() getAmmeterSettingsCommand {
	c := getAmmeterSettingsCommand{}
	c.Type = command.GetAmmeterSettings

	return c
}

func parseResponse(response string) (currentScaleFactor int, currentOffset int, err error) {
	split := strings.Split(response, " ")

	if currentScaleFactor, err = util.ParseInt(split[1]); err != nil {
		return
	}

	if currentOffset, err = util.ParseInt(split[2]); err != nil {
		return
	}

	return
}
