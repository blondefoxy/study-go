package facade

import "fmt"

type healthWorker interface {
	Check() bool
}

type recommender interface {
	Write() string
}

// GenomService service for calculating gemon data
type GenomService struct {
	geneData   string
	health     healthWorker
	recommends recommender
}

// NewGenomService creates new genom service
func NewGenomService(data string, health healthWorker, recommend recommender) *GenomService {
	return &GenomService{
		geneData:   data,
		health:     health,
		recommends: recommend,
	}
}

// Calculate make all calculations to check health and give recommendations
func (g *GenomService) Calculate() error {
	g.health.Check()
	fmt.Println(g.recommends.Write())

	return nil
}
