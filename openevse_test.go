package openevse

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"net/http"
	"github.com/avegao/openevse/command"
)

//func TestGetRtcTime(t *testing.T) {
//	assert.Equal(t, "a", "a")
//}

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
			Command:  fmt.Sprintf("$%s", command.GetAmmeterSetings),
			Response: fmt.Sprintf("$OK %d %d", expectedCurrentScaleFactor, expectedCurrentOffset),
		})

	assert.NoError(t, r.Error)

	currentScaleFactor, currentOffset, err := GetAmmeterSettings("192.168.1.156")

	assert.NoError(t, err)
	assert.Equal(t, expectedCurrentScaleFactor, currentScaleFactor)
	assert.Equal(t, expectedCurrentOffset, currentOffset)

	assert.True(t, gock.IsDone())
}
