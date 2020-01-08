package facade

import "fmt"

type healthWorker interface {
	Check() bool
}

type recommender interface {
	Write() string
}

// GenomServicer ...
type GenomServicer interface {
	Calculate() error
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

// NewGenomServicer creates new genom service
func NewGenomServicer(data string, health healthWorker, recommend recommender) GenomServicer {
	return &genomService{
		geneData:   data,
		health:     health,
		recommends: recommend,
	}
}
