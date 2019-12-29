package facade

import "fmt"

// GenomService ...
type GenomService interface {
	Calculate() error
}

type healthWorker interface {
	Check() bool
}

type recommender interface {
	Write() string
}

type genomService struct {
	geneData   string
	health     healthWorker
	recommends recommender
}

// Calculate make all calculations to check health and give recommendations
func (g *genomService) Calculate() error {
	g.health.Check()
	fmt.Println(g.recommends.Write())
	return nil
}

// NewGenomService creates new genom service
func NewGenomService(data string, health healthWorker, recommend recommender) GenomService {
	return &genomService{
		geneData:   data,
		health:     health,
		recommends: recommend,
	}
}
