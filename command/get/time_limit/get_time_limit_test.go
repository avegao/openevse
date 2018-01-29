package getTimeLimit

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"github.com/avegao/openevse/command"
)

func TestGetTimeLimitCommand_Run(t *testing.T) {
	defer gock.Disable()

	tests := []struct {
		expected int
	}{
		{0}, {10}, {20}, {30}, {40}, {50}, {60},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.expected), func(t *testing.T) {
			r := gock.New("http://(.*)").
				Reply(http.StatusOK).
				JSON(
				command.Response{
					Command:  fmt.Sprintf("$%s", command.GetTimeLimit),
					Response: fmt.Sprintf("$OK %d", test.expected),
				})

			assert.NoError(t, r.Error)

			timeLimit, err := New().Run("192.168.1.156")

			assert.NoError(t, err)
			assert.Equal(t, test.expected, timeLimit)
		})
	}

	assert.True(t, gock.IsDone())
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, New())
}

