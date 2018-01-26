package getFaultCounters

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
	"fmt"
	"github.com/avegao/openevse/command"
	"net/http"
)

func TestGetFaultCounters_Run(t *testing.T) {
	defer gock.Disable()

	r := gock.New("http://(.*)").
		Reply(http.StatusOK).
		JSON(
		command.Response{
			Command:  fmt.Sprintf("$%s", command.GetFaultCounters),
			Response: fmt.Sprint("$OK f a b"),
		})

	assert.NoError(t, r.Error)

	gfdi, noGround, stuckRelay, err := New().Run("192.168.1.156")

	assert.NoError(t, err)

	assert.Equal(t, 15, gfdi)
	assert.Equal(t, 10, noGround)
	assert.Equal(t, 11, stuckRelay)
	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New())
}
