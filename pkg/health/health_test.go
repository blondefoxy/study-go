package health

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	healthBehaviorTestName = "check health behavior"
)

func TestHealth(t *testing.T) {
	t.Run(healthBehaviorTestName, func(t *testing.T) {
		hlth := &health{}
		require.True(t, hlth.Check())
	})
}
