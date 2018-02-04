package getAmmeterSettings

import (
	"github.com/avegao/openevse/command"
	"strings"
	"github.com/avegao/openevse/util"
)

type getAmmeterSettingsCommandInterface interface {
	Run() (currentScaleFactor int, currentOffset int, err error)
	parseResponse(response string) (currentScaleFactor int, currentOffset int, err error)
}

type getAmmeterSettingsCommand struct {
	getAmmeterSettingsCommandInterface
	command.Command
}

func (c getAmmeterSettingsCommand) Run() (currentScaleFactor int, currentOffset int, err error) {
	response, err := c.SendRequest()

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getAmmeterSettingsCommand) parseResponse(response string) (currentScaleFactor int, currentOffset int, err error) {
	split := strings.Split(response, " ")

	if currentScaleFactor, err = util.ParseInt(split[1]); err != nil {
		return
	}

	if currentOffset, err = util.ParseInt(split[2]); err != nil {
		return
	}

	return
}

func New(host string) getAmmeterSettingsCommand {
	c := getAmmeterSettingsCommand{}
	c.Host = host
	c.Type = command.GetAmmeterSettings

	return c
}
