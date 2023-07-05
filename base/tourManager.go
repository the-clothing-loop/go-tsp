package base

// ToueManager : Contains list of of cities to be visited
type TourManager struct {
	destCities []City
}

// NewTourManager : Initialize TourManager
func (a *TourManager) NewTourManager() {
	a.destCities = []City{}
}

func (a *TourManager) AddCity(c City) {
	a.destCities = append(a.destCities, c)
}

func (a *TourManager) GetCity(i int) City {
	return a.destCities[i]
}

func (a *TourManager) NumberOfCities() int {
	return len(a.destCities)
}

func (a *TourManager) GetCities() []City {
	return a.destCities
}

func (tm *TourManager) GetDistanceMatrix() [][]float64 {
	numberOfCities := tm.NumberOfCities()
	distances := make([][]float64, numberOfCities)
	for i := 0; i < numberOfCities; i++ {
		distances[i] = make([]float64, numberOfCities)
		for j := 0; j < numberOfCities; j++ {
			aCity := tm.GetCity(i)
			bCity := tm.GetCity(j)

			distances[i][j] = aCity.DistanceTo(bCity)
		}
	}
	return distances
}

func InitRandomTourManager(numberOfCities int) TourManager {
	tm := TourManager{}
	tm.NewTourManager()

	cities := initRandomCities(numberOfCities)
	for _, v := range cities {
		tm.AddCity(v)
	}
	return tm
}

func initRandomCities(cityCount int) []City {
	cities := make([]City, 0, cityCount)
	for i := 0; i < cityCount; i++ {
		cities = append(cities, GenerateRandomCity())
	}
	return cities
}
