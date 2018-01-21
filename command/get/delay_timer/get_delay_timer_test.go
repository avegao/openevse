package getDelayTimer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"github.com/avegao/openevse/command"
	"net/http"
)

func TestGetDelayTimer_Run(t *testing.T) {
	const expectedStartTime = "00:00:00"
	const expectedEndTime = "08:00:00"

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetDelayTimer),
			Response: fmt.Sprintf("$OK %d %d %d %d", 0, 0, 8, 0),
		})

	assert.NoError(t, r.Error)

	startTime, endTime, err := New().Run("192.168.1.156")

	assert.NoError(t, err)
	assert.NotEmpty(t, startTime)
	assert.Equal(t, expectedStartTime, startTime)
	assert.NotEmpty(t, endTime)
	assert.Equal(t, expectedEndTime, endTime)

	assert.True(t, gock.IsDone())
}
