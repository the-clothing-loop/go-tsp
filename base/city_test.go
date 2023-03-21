package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-clothing-loop/go-tsp/base"
)

func TestShuffleCities(t *testing.T) {
	cities := []base.City{
		base.GenerateCity(1, 1, 1),
		base.GenerateCity(1, 2, 2),
		base.GenerateCity(1, 3, 3),
		base.GenerateCity(1, 4, 4),
		base.GenerateCity(1, 5, 5),
	}

	randCities := base.ShuffleCities(cities)

	assert.Contains(t, randCities, cities[0])
	assert.Contains(t, randCities, cities[1])
	assert.Contains(t, randCities, cities[2])
	assert.Contains(t, randCities, cities[3])
	assert.Contains(t, randCities, cities[4])

	assert.NotEqual(t, cities, randCities)
	assert.Equal(t, 5, len(randCities))
}
