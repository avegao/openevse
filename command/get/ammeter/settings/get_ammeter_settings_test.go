package getAmmeterSettings

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"github.com/avegao/openevse/command"
	"net/http"
)

func TestGetAmmeterSettingsCommand_Run(t *testing.T) {
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

	currentScaleFactor, currentOffset, err := New("192.168.1.156").Run()

	assert.NoError(t, err)
	assert.Equal(t, expectedCurrentScaleFactor, currentScaleFactor)
	assert.Equal(t, expectedCurrentOffset, currentOffset)

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New("192.168.1.156"))
}
