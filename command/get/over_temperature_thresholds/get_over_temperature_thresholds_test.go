package getOverTemperatureThresholds

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"net/http"
	"github.com/avegao/openevse/command"
)

func TestGetOverTemperatureThresholds_Run(t *testing.T) {
	const (
		expectedAmbient float32 = 89.5
		expectedIr float32 = 56.5
	)

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetOverTemperatureThresholds),
			Response: fmt.Sprint("$OK 895 565"),
		})

	assert.NoError(t, r.Error)

	ambient, ir, err := New("192.168.1.156").Run()

	assert.NoError(t, err)
	assert.Equal(t, expectedAmbient, ambient)
	assert.Equal(t, expectedIr, ir)
	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New("192.168.1.156"))
}
