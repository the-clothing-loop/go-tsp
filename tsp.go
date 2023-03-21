package tsp

import (
	"github.com/the-clothing-loop/go-tsp/base"
	ga "github.com/the-clothing-loop/go-tsp/geneticAlgorithm"
)

// tspGA : Travelling sales person with genetic algorithm
// input :- Cities, Number of generations

func TspGA(cities []base.City, gen int) (*base.Tour, *base.Population) {
	tm := &base.TourManager{}
	for _, city := range cities {
		tm.AddCity(city)
	}

	p := base.Population{}
	p.InitPopulation(200, *tm)

	// Evolve population "gen" number of times
	for i := 1; i < gen+1; i++ {
		p = ga.EvolvePopulation(p)
	}
	return p.GetFittest(), &p
}
