package facade

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	serviceBehaviorSuccess = "check service behavior"
	inputData              = "foo"
	recommenderAnswer      = "test passed successful"
)

func TestGenomService(t *testing.T) {
	t.Run(serviceBehaviorSuccess, func(t *testing.T) {
		genom := NewGenomService(inputData, newMockHealth(), newMockRecommender())
		require.NoError(t, genom.Calculate())
	})
}

type mockHealth struct {
}

func (h *mockHealth) Check() bool {
	return true
}

func newMockHealth() healthWorker {
	return &mockHealth{}
}

type mockRecommender struct {
}

func (r *mockRecommender) Write() string {
	return recommenderAnswer
}

func newMockRecommender() recommender {
	return &mockRecommender{}
}
