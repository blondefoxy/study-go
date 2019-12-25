package recommends

// Recommender ...
type Recommender interface {
	Write() string
}

type recommend struct {
}

// NewRecommender ...
func NewRecommender() Recommender {
	return &recommend{}
}

// Write - make a recommendation
func (r *recommend) Write() string {
	return "recommend training"
}
