package health

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHealth(t *testing.T) {
	t.Run("check health behavior", func(t *testing.T) {
		hlth := &health{}
		require.True(t, hlth.Check())
	})
}
