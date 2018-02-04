package getVoltmeterSettings

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"github.com/avegao/openevse/command"
	"net/http"
)

func TestGetVoltmeterSettings_Run(t *testing.T) {
	const (
		calefactorExpected = 100
		offsetExpected     = 100
	)

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetVoltmeterSettings),
			Response: fmt.Sprintf("$OK %d %d", calefactorExpected, offsetExpected),
		})

	assert.NoError(t, r.Error)

	calefactor, offset, err := New("192.168.1.156").Run()

	assert.NoError(t, err)
	assert.Equal(t, calefactorExpected, calefactor)
	assert.Equal(t, offsetExpected, offset)

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New("192.168.1.156"))
}
