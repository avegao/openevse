package getSettings

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"github.com/avegao/openevse/command"
	"net/http"
)

func TestGetSettingsCommand_Run(t *testing.T) {
	const (
		expectedAmperes = 16
		expectedFlags   = 33
	)

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetSettings),
			Response: "$OK 16 0021^16",
		})

	assert.NoError(t, r.Error)

	amperes, flags, err := New("192.168.1.156").Run()

	assert.NoError(t, err)
	assert.Equal(t, expectedAmperes, amperes)
	assert.Equal(t, expectedFlags, flags)
	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New("192.168.1.156"))
}
