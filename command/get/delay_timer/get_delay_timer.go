package getDelayTimer

import (
	"strings"
	"errors"
	"github.com/avegao/openevse/command"
	"fmt"
)

type getDelayTimerInterface interface {
	Run(response string) (startTime string, endTime string, err error)
	parseResponse(response string) (startTime string, endTime string, err error)
}

type getDelayTimer struct {
	getDelayTimerInterface
	command.Command
}

func (c getDelayTimer) Run(host string) (startTime string, endTime string, err error) {
	c.Type = command.GetDelayTimer

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getDelayTimer) parseResponse(response string) (startTime string, endTime string, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		startTime, endTime = parseResponseFromSplit(split)
	case command.FailureResponse:
		err = errors.New("openevse - invalid request")
	default:
		err = errors.New("openevse - invalid request - unknown error")
	}

	return
}

func New() getDelayTimer {
	c := new(getDelayTimer)
	c.Type = command.GetDelayTimer

	return *c
}

func parseResponseFromSplit(split []string) (startTime string, endTime string) {
	startHour := parseHour(split[1])
	startMinutes := parseMinute(split[2])
	endHour := parseHour(split[3])
	endMinutes := parseMinute(split[4])

	startTime = fmt.Sprintf("%s:%s:00", startHour, startMinutes)
	endTime = fmt.Sprintf("%s:%s:00", endHour, endMinutes)

	return
}

func parseHour(hour string) string {
	if len(hour) == 1 {
		hour = "0" + hour
	}

	return hour
}

func parseMinute(minutes string) string {
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}

	return minutes
}
