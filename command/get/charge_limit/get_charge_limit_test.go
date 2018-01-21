package getChargeLimit

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"net/http"
	"github.com/avegao/openevse/command"
)

func TestGetChargeLimit_Run(t *testing.T) {
	const kwhExpected = 25

	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetChargeLimit),
			Response: fmt.Sprintf("$OK %d", kwhExpected),
		})

	assert.NoError(t, r.Error)

	kwh, err := New().Run("192.168.1.156")

	assert.NoError(t, err)
	assert.NotZero(t, kwh)
	assert.Equal(t, kwhExpected, kwh)

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New())
}
