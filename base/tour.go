package base

import (
	"strconv"
)

type Tour struct {
	tourCities []City
	fitness    float64
	distance   float64
}

// InitTour : Initialize tour with cities arranged randomly
func (a *Tour) InitTour(numberOfCities int) {
	a.tourCities = make([]City, numberOfCities)
}

// InitTourCities
func (a *Tour) InitTourCities(tm TourManager) {
	a.InitTour(tm.NumberOfCities())
	// Add all destination cities from TourManager to Tour
	for i := 0; i < tm.NumberOfCities(); i++ {
		a.SetCity(i, tm.GetCity(i))
	}
	// Shuffle cities in tour
	a.tourCities = ShuffleCities(a.tourCities)
}

// GetCity : Get city based on position in slice
func (a *Tour) GetCity(tourPosition int) City {
	return a.tourCities[tourPosition]
}

// SetCity : Set position of city in tour slice
func (a *Tour) SetCity(tourPosition int, c City) {
	a.tourCities[tourPosition] = c
	// Reset fitness if tour have been altered
	a.fitness = 0
	a.distance = 0
}

func (a *Tour) ResetFitnessDistance() {
	a.fitness = 0
	a.distance = 0
}

func (a *Tour) TourSize() int {
	return len(a.tourCities)
}

// TourDistance : Calculates total distance traveled for this tour
func (a *Tour) TourDistance() float64 {
	if a.distance == 0 {
		td := float64(0)
		for i := 0; i < a.TourSize(); i++ {
			fromC := a.GetCity(i)
			var destC City
			if i+1 < a.TourSize() {
				destC = a.GetCity(i + 1)
			} else {
				destC = a.GetCity(0)
			}
			td += fromC.DistanceTo(destC)
		}
		a.distance = td
	}
	return a.distance
}

func (a *Tour) Fitness() float64 {
	if a.fitness == 0 {
		a.fitness = 1 / a.TourDistance()
	}
	return a.fitness
}

func (a *Tour) ContainCity(c City) bool {
	for _, cs := range a.tourCities {
		if cs == c {
			return true
		}
	}
	return false
}

func (a Tour) String() string {
	s := "|"
	for i, c := range a.tourCities {
		s += strconv.Itoa(i) + c.String() + "|"
	}
	return s
}

// Return an ordered slice of the tour
// See this as a route listed by the IDs
// The firstID argument will be the first ID in the result
// The firstHalfID argument will be in the first half of the result
func (a Tour) RouteByIDs(firstID, firstHalfID uint) []uint {
	route := make([]uint, 0, len(a.tourCities))
	for _, c := range a.tourCities {
		route = append(route, c.id)
	}

	// turn route to set firstID to index 0
	return reorder(route, firstID, firstHalfID)
}

// reorder `arr` to return a slice
//
// `a` is the first value inside the return slice
// `b` is a value somewhere inside the first half of the return slice
func reorder(arr []uint, a, b uint) []uint {
	if len(arr) < 2 {
		return arr
	}

	aIndex := -1
	for i, v := range arr {
		if v == a {
			aIndex = i
			break
		}
	}

	first := arr[aIndex]
	next := []uint{}
	if aIndex < len(arr) {
		next = arr[aIndex+1:]
	}
	if aIndex > 0 {
		next = append(next, arr[:aIndex]...)
	}

	isBInFirstHalf := false
	for _, v := range next[:len(next)/2] {
		if v == b {
			isBInFirstHalf = true
		}
	}
	if !isBInFirstHalf {
		reverse(next)
	}

	arr = append([]uint{first}, next...)
	return arr
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
