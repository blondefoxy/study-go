package facade

import (
	"testing"

	"github.com/blondefoxy/study-go/pkg/health"
	"github.com/blondefoxy/study-go/pkg/recommends"
	"github.com/stretchr/testify/require"
)

func TestGenomService(t *testing.T) {
	t.Run("check service behavior", func(t *testing.T) {
		genom := &GenomService{
			geneData:   "",
			health:     health.NewHealth(),
			recommends: recommends.NewRecommender(),
		}
		err := genom.Calculate()
		require.NoError(t, err)
	})
}
