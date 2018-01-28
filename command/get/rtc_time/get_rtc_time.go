package getRtcTime

import (
	"time"
	"github.com/avegao/openevse_http/command"
	"strings"
	"errors"
	"fmt"
	"github.com/avegao/openevse_http/util"
)

type getRtcTimeCommandInterface interface {
	Run(response string) (rtcTime time.Time, err error)
	parseResponse(response string) (rtcTime time.Time, err error)
}

type getRtcTimeCommand struct {
	getRtcTimeCommandInterface
	command.Command
}

func (c getRtcTimeCommand) Run(host string) (rtcTime time.Time, err error) {
	c.Type = command.GetRtcTime

	response, err := c.SendRequest(host)

	if err != nil {
		return
	}

	return c.parseResponse(response.Response)
}

func (c getRtcTimeCommand) parseResponse(response string) (rtcTime time.Time, err error) {
	split := strings.Split(response, " ")

	switch split[0] {
	case command.SuccessResponse:
		rtcTime, err = parseTimeFromResponseSplit(split)
	case command.FailureResponse:
		err = errors.New("invalid request")
	default:
		err = errors.New("invalid request - unknown error")
	}

	return
}

func New() getRtcTimeCommand {
	c := new(getRtcTimeCommand)
	c.Type = command.GetRtcTime

	return *c
}

func parseTimeFromResponseSplit(split []string) (rtcTime time.Time, err error) {
	var year, month, day, hour, minute, second string

	year = "20" + split[1]

	if month, err = parseTimeElement(split[2], true); err != nil {
		return
	}

	if day, err = parseTimeElement(split[3], false); err != nil {
		return
	}

	if hour, err = parseTimeElement(split[4], false); err != nil {
		return
	}

	if minute, err = parseTimeElement(split[5], false); err != nil {
		return
	}

	if second, err = parseTimeElement(split[6], false); err != nil {
		return
	}

	return time.Parse(time.RFC3339, fmt.Sprintf("%s-%s-%sT%s:%s:%sZ", year, month, day, hour, minute, second))
}

func parseTimeElement(element string, isMonth bool) (elementString string, err error) {
	elementInt, err := util.ParseInt(element)

	if err != nil {
		return
	}

	if isMonth {
		elementInt++
	}

	elementString = fmt.Sprintf("%d", elementInt)

	if elementInt < 10 {
		elementString = "0" + elementString
	}

	return
}
