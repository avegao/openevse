package openevse

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"net/http"
	"github.com/avegao/openevse/command"
	"time"
)

func TestGetRtcTime(t *testing.T) {
	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetRtcTime),
			Response: "$OK 18 0 7 16 24 9^16",
		})

	assert.NoError(t, r.Error)

	rtcTime, err := GetRtcTime("192.168.1.156")

	assert.NoError(t, err)
	assert.Equal(t, rtcTime, time.Date(2018, 1, 7, 16, 24, 9, 0, time.UTC))

	now := time.Now()
	// this is to remove nano seconds
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.UTC)
	responseTime := fmt.Sprintf("$OK %d %d %d %d %d %d^16", now.Year() - 2000, now.Month() - 1, now.Day(), now.Hour(), now.Minute(), now.Second())

	r = gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetRtcTime),
			Response: responseTime,
		})

	assert.NoError(t, r.Error)

	rtcTime, err = GetRtcTime("192.168.1.156")

	assert.NoError(t, err)
	assert.Equal(t, rtcTime, now)

	assert.True(t, gock.IsDone())
}

func TestGetVersion(t *testing.T) {
	const firmwareVersionExpected = "4.8.0"
	const protocolVersionExpected = "3.0.1"

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetRtcTime),
			Response: fmt.Sprintf("$OK %s %s", firmwareVersionExpected, protocolVersionExpected),
		})

	assert.NoError(t, r.Error)

	GetVersion("192.168.1.156")
	assert.Equal(t, "a", "a")
}

func TestGetAmmeterSettings(t *testing.T) {
	const expectedCurrentScaleFactor = 220
	const expectedCurrentOffset = 0

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetAmmeterSettings),
			Response: fmt.Sprintf("$OK %d %d", expectedCurrentScaleFactor, expectedCurrentOffset),
		})

	assert.NoError(t, r.Error)

	currentScaleFactor, currentOffset, err := GetAmmeterSettings("192.168.1.156")

	assert.NoError(t, err)
	assert.Equal(t, expectedCurrentScaleFactor, currentScaleFactor)
	assert.Equal(t, expectedCurrentOffset, currentOffset)

	assert.True(t, gock.IsDone())
}
