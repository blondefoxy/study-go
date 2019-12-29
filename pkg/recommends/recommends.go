package recommends

const training = "recommend training"

// Recommender ...
type Recommender interface {
	Write() string
}

type recommend struct {
}

// Write - make a recommendation
func (r *recommend) Write() string {
	return training
}

// NewRecommender ...
func NewRecommender() Recommender {
	return &recommend{}
}
