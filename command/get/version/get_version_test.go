package getVersion

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"net/http"
	"github.com/avegao/openevse/command"
)

func TestGetVersionCommand_Run(t *testing.T) {
	const firmwareVersionExpected = "4.8.0"
	const protocolVersionExpected = "3.0.1"

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetVersion),
			Response: fmt.Sprintf("$OK %s %s", firmwareVersionExpected, protocolVersionExpected),
		})

	assert.NoError(t, r.Error)

	firmwareVersion, protocolVersion, err := New().Run("192.168.1.156")

	assert.NoError(t, err)
	assert.NotEmpty(t, firmwareVersion)
	assert.Equal(t, firmwareVersionExpected, firmwareVersion)
	assert.NotEmpty(t, protocolVersion)
	assert.Equal(t, protocolVersionExpected, protocolVersion)

	assert.True(t, gock.IsDone())
}
