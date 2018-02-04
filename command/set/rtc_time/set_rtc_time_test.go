package setRtcTime

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"github.com/avegao/openevse/command"
	"net/http"
)

func TestGetRtcTimeCommand_Run(t *testing.T) {
	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.SetRtcTime),
			Response: "$OK 18 0 7 16 24 9^16",
		})

	assert.NoError(t, r.Error)

	err := New("192.168.1.156").Run(time.Now())

	assert.NoError(t, err)
	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New("192.168.1.156"))
}
