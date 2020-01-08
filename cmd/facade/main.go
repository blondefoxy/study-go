package main

import (
	"log"

	"github.com/blondefoxy/study-go/pkg/facade"
	"github.com/blondefoxy/study-go/pkg/health"
	"github.com/blondefoxy/study-go/pkg/recommends"
)

const data = "1qaz"

func main() {
	err := facade.NewGenomServicer(data,
		health.NewHealther(),
		recommends.NewRecommender(),
	).Calculate()
	if err != nil {
		log.Printf("error occured: %v", err)
	}

	log.Println("OK")
}
