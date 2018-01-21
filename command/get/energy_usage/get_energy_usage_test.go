package getEnergyUsage

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"net/http"
	"github.com/avegao/openevse/command"
)

func TestGetEnergyUsage_Run(t *testing.T) {
	const expectedWhInSession = 35365
	const expectedWhAccumulated = 382

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetEnergyUsage),
			Response: fmt.Sprintf("$OK %d %d", expectedWhInSession*3600, expectedWhAccumulated*1000),
		})

	assert.NoError(t, r.Error)

	whInSession, whAccumulated, err := New().Run("192.168.1.156")

	assert.NoError(t, err)
	assert.NotZero(t, whInSession)
	assert.Equal(t, expectedWhInSession, whInSession)
	assert.NotZero(t, whAccumulated)
	assert.Equal(t, expectedWhAccumulated, whAccumulated)

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New())
}
