package base

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTourDistance(t *testing.T) {
	fmt.Println("Traveling sales person - Standard Random")
	// Init TourManager
	tm := TourManager{}
	tm.NewTourManager()

	// Generate Cities
	cities := *initializeSampleCities()

	// Add cities to TourManager
	for _, v := range cities {
		tm.AddCity(v)
	}

	// Init population
	p := Population{}
	pSize := 50
	p.InitPopulation(pSize, tm)
	fmt.Println("Find........")
	allDistances := make([]float64, 0, 20)
	shortest := p.GetTour(0).TourDistance()
	for i := 0; i < pSize; i++ {
		d := p.GetTour(i).TourDistance()
		allDistances = append(allDistances, d)
		if shortest > d {
			shortest = d
		}
	}

	fmt.Println("Initial best distance: ", p.GetFittest().TourDistance())
	assert.Equalf(t, shortest, p.GetFittest().TourDistance(), "find the shortest distance from %+v", allDistances)
}

func cityListContain() {
	c1 := City{}
	c1.SetLocation(10, 20)
	c2 := City{}
	c2.SetLocation(30, 40)
	c3 := City{}
	c3.SetLocation(50, 60)
	c4 := City{}
	c4.SetLocation(30, 40)

	cslice := make([]City, 0, 20)

	cslice = append(cslice, c1)
	cslice = append(cslice, c2)
	cslice = append(cslice, c3)

	fmt.Println(cslice)
	fmt.Println(cslice[0])

	for _, c := range cslice {
		if c == c4 {
			fmt.Println("found same", c)
		}
		if reflect.DeepEqual(c, c4) {
			fmt.Println("deep equal true", c)
		}
	}

}

func initializeSampleCities() *[]City {
	cities := make([]City, 0, 20)
	// Sample
	cities = append(cities, GenerateCity(60, 200, 1))
	cities = append(cities, GenerateCity(180, 200, 2))
	cities = append(cities, GenerateCity(80, 180, 3))
	cities = append(cities, GenerateCity(140, 180, 4))
	cities = append(cities, GenerateCity(20, 160, 5))
	cities = append(cities, GenerateCity(100, 160, 6))
	cities = append(cities, GenerateCity(200, 160, 7))
	cities = append(cities, GenerateCity(140, 140, 8))
	cities = append(cities, GenerateCity(40, 120, 9))
	cities = append(cities, GenerateCity(100, 120, 10))
	cities = append(cities, GenerateCity(180, 100, 11))
	cities = append(cities, GenerateCity(60, 80, 12))
	cities = append(cities, GenerateCity(120, 80, 13))
	cities = append(cities, GenerateCity(180, 60, 14))
	cities = append(cities, GenerateCity(20, 40, 15))
	cities = append(cities, GenerateCity(100, 40, 16))
	cities = append(cities, GenerateCity(200, 40, 17))
	cities = append(cities, GenerateCity(20, 20, 18))
	cities = append(cities, GenerateCity(60, 20, 19))
	cities = append(cities, GenerateCity(160, 20, 20))

	// Sample using random seed
	// Completed testing
	return &cities
}

func TestReorder(t *testing.T) {
	expected := []uint{1, 2, 3, 4, 5}
	assert.Equal(t, expected, reorder([]uint{1, 2, 3, 4, 5}, 1, 2))
	assert.Equal(t, expected, reorder([]uint{5, 1, 2, 3, 4}, 1, 2))
	assert.Equal(t, expected, reorder([]uint{4, 5, 1, 2, 3}, 1, 2))
	assert.Equal(t, expected, reorder([]uint{3, 4, 5, 1, 2}, 1, 2))
	assert.Equal(t, expected, reorder([]uint{2, 3, 4, 5, 1}, 1, 2))
}

func TestReorderArrOfOne(t *testing.T) {
	assert.Equal(t, []uint{1}, reorder([]uint{1}, 1, 1))
}
