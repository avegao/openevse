package getChargingCurrentAndVoltage

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"github.com/avegao/openevse/command"
	"net/http"
)

func TestGetChargeLimit_Run(t *testing.T) {
	const milliAmpsExpected = 16000
	const milliVoltsExpected = 230

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetChargingCurrentAndVoltage),
			Response: fmt.Sprintf("$OK %d %d", milliAmpsExpected, milliVoltsExpected),
		})

	assert.NoError(t, r.Error)

	milliAmps, milliVolts, err := New().Run("192.168.1.156")

	assert.NoError(t, err)
	assert.NotZero(t, milliAmps)
	assert.Equal(t, milliAmpsExpected, milliAmps)
	assert.NotZero(t, milliVolts)
	assert.Equal(t, milliVoltsExpected, milliVolts)

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New())
}
