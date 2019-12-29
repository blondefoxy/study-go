package recommends

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	recommendsTestName       = "check recommends behavior"
	recommendsAnswerExpected = "recommend training"
)

func TestHealth(t *testing.T) {
	t.Run(recommendsTestName, func(t *testing.T) {
		rec := &recommend{}
		result := rec.Write()
		require.NotEmpty(t, result)
		require.Exactly(t, recommendsAnswerExpected, result)
	})
}
