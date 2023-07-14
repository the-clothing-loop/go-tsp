package tsp

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-clothing-loop/go-tsp/algorithms/mst"
	"github.com/the-clothing-loop/go-tsp/base"
)

// * 1 * * * 5 * * * *
// * * * * * * * * * *
// * * * * * * * * * *
// * * * * * * * * * *
// * * * * * * * * 4 *
// 2 * * * * * * * * *
// * * * * * * * * * *
// * * * * * * * * * *
// * * * * * * * 3 * *
// * * * * * * * * * *

func TestTsp(t *testing.T) {
	cities := []base.City{
		base.GenerateCity(2, 10, 1),
		base.GenerateCity(1, 5, 2),
		base.GenerateCity(8, 2, 3),
		base.GenerateCity(9, 6, 4),
		base.GenerateCity(6, 10, 5),
	}

	tour, _ := TspGA(cities, 20)

	route := tour.RouteByIDs(1, 2)

	if !assert.ObjectsAreEqual([]uint{1, 2, 3, 4, 5}, route) {
		assert.Fail(t, "Incorrect route: %v", route)
	}
}

func BenchmarkTspGenetic(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		cities := []base.City{
			base.GenerateCity(2, 10, 1),
			base.GenerateCity(1, 5, 2),
			base.GenerateCity(8, 2, 3),
			base.GenerateCity(9, 6, 4),
			base.GenerateCity(6, 10, 5),
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			tour, _ := TspGA(cities, 20)

			route := tour.RouteByIDs(1, 2)
			if !assert.ObjectsAreEqual([]uint{1, 2, 3, 4, 5}, route) {
				assert.Fail(b, "Incorrect route: %v", route)
			}
		}()
	}
	wg.Wait()
}

func BenchmarkTspGeneticVsMST(b *testing.B) {

	// initialize multiple tourManager with different number of cities
	numberOfTourManager := 10
	tableTest := make([]struct {
		cities int
		tm     base.TourManager
	}, numberOfTourManager)

	numberOfCities := 50

	for i := 0; i < numberOfTourManager; i++ {
		tm := base.InitRandomTourManager(numberOfCities)
		tableTest[i] = struct {
			cities int
			tm     base.TourManager
		}{
			cities: numberOfCities,
			tm:     tm,
		}
		numberOfCities += 50
	}

	b.Run("Run GA", func(b *testing.B) {
		for _, v := range tableTest {
			b.Run(fmt.Sprintf("number_cities_%d", v.cities), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					TspGA(v.tm.GetCities(), 50)
				}
			})
		}
	})

	b.Run("Run MST", func(b *testing.B) {
		for _, v := range tableTest {
			b.Run(fmt.Sprintf("number_cities_%d", v.cities), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					mst.MST(v.tm.GetDistanceMatrix())
				}
			})
		}
	})

}
