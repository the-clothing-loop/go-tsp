package tsp

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
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

	route := tour.RouteByIDs(1)

	if !(assert.ObjectsAreEqual([]uint{1, 2, 3, 4, 5}, route) || assert.ObjectsAreEqual([]uint{1, 5, 4, 3, 2}, route)) {
		assert.Fail(t, "Incorrect route: %v", route)
	}
}

func BenchmarkTestTsp(b *testing.B) {
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

			route := tour.RouteByIDs(1)
			if !(assert.ObjectsAreEqual([]uint{1, 2, 3, 4, 5}, route) || assert.ObjectsAreEqual([]uint{1, 5, 4, 3, 2}, route)) {
				assert.Fail(b, "Incorrect route: %v", route)
			}
		}()

	}
	wg.Wait()
}
