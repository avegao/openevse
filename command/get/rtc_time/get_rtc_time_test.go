package getRtcTime

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/stretchr/testify/assert"
	"fmt"
	"time"
	"github.com/avegao/openevse_http/command"
	"net/http"
)

func TestGetRtcTimeCommand_Run(t *testing.T) {
	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetRtcTime),
			Response: "$OK 18 0 7 16 24 9^16",
		})

	assert.NoError(t, r.Error)

	rtcTime, err := New().Run("foo.com")

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

	rtcTime, err = New().Run("foo.com")

	assert.NoError(t, err)
	assert.Equal(t, rtcTime, now)

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New())
}
