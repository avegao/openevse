package currentCapacityRangeInAmps

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"net/http"
	"github.com/avegao/openevse/command"
)

func TestGetCurrentCapacityRangeInAmps_Run(t *testing.T) {
	const minAmpsExpected = 10
	const maxAmpsExpected = 80

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetCurrentCapacityRangeInAmps),
			Response: fmt.Sprintf("$OK %d %d", minAmpsExpected, maxAmpsExpected),
		})

	assert.NoError(t, r.Error)

	minAmps, maxAmps, err := New().Run("192.168.1.156")

	assert.NoError(t, err)

	assert.NotZero(t, minAmps)
	assert.Equal(t, minAmpsExpected, minAmps)
	assert.NotZero(t, maxAmps)
	assert.Equal(t, maxAmpsExpected, maxAmps)

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New())
}
