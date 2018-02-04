package setRtcTime

import (
	"time"
	"github.com/avegao/openevse/command"
	"fmt"
	"strings"
	"errors"
)

type setRtcTimeCommandInterface interface {
	Run(host string, rtcTime time.Time) (err error)
}

type setRtcTimeCommand struct {
	setRtcTimeCommandInterface
	command.Command
}

func (c setRtcTimeCommand) Run(host string, rtcTime time.Time) (err error) {
	c.Type = command.SetRtcTime

	year := []byte(fmt.Sprintf("%d", rtcTime.Year()))
	month := rtcTime.Month() - 1

	timeString := fmt.Sprintf("%s%s %d %d %d %d %d",
		string(year[2]),
		string(year[3]),
		month,
		rtcTime.Day(),
		rtcTime.Hour(),
		rtcTime.Minute(),
		rtcTime.Second(),
	)

	response, err := c.SendRequest(host, timeString)

	if err != nil {
		return
	}

	err = c.parseResponse(response.Response)

	return
}

func (c setRtcTimeCommand) parseResponse(response string) (err error) {
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

func New() setRtcTimeCommand {
	c := new(setRtcTimeCommand)
	c.Type = command.SetRtcTime

	return *c
}
