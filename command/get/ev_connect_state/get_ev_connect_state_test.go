package getEvConnectState

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"github.com/avegao/openevse/command"
	"github.com/avegao/openevse/command/ev_connect_state"
)

func TestGetEvConnectState_Run(t *testing.T) {
	tests := []evConnectState.EvConnectState{
		evConnectState.NotConnected,
		evConnectState.Connected,
		evConnectState.Unknown,
	}

	defer gock.Disable()

	for _, expected := range tests {
		r := gock.New("http://(.*)").
			Reply(http.StatusOK).
			JSON(
			command.Response{
				Command:  fmt.Sprintf("$%s", command.GetEvConnectState),
				Response: fmt.Sprintf("$OK %d", expected),
			})

		assert.NoError(t, r.Error)

		state, err := New("192.168.1.156").Run()

		assert.NoError(t, err)
		assert.Equal(t, expected, state)
	}

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New("192.168.1.156"))
}
