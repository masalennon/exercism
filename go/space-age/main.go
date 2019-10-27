package space

func main() {

}

func Age(seconds float64, planet Planet) float64 {
	yearOnEarth := calculateYearOnEarth(seconds)
	switch planet {
	case "Earth":
		return yearOnEarth
	case "Mercury":
		return yearOnEarth / 0.2408467
	case "Venus":
		return yearOnEarth / 0.61519726
	case "Mars":
		return yearOnEarth / 1.8808158
	case "Jupiter":
		return yearOnEarth / 11.862615
	case "Saturn":
		return yearOnEarth / 29.447498
	case "Uranus":
		return yearOnEarth / 84.016846
	case "Neptune":
		return yearOnEarth / 164.79132
	default:
		return yearOnEarth
	}
}

func calculateYearOnEarth(seconds float64) float64 {
	return float64(seconds / 31557600)
}


type Planet string