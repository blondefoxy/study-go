package recommends

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	t.Run("check recommends behavior", func(t *testing.T) {
		rec := &recommend{}
		result := rec.Write()
		require.NotEmpty(t, result)
		require.Exactly(t, "recommend training", result)
	})
}
