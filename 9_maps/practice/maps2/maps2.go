package main

import "fmt"

func addFilm(newFilm, catalog map[string]float64) {
	for name, rating := range newFilm {
		_, ok := catalog[name]
		if ok {
			continue
		}
		catalog[name] = rating

	}
}

func topFilms(catalog map[string]float64, min float64) []string {
	result := []string{}

	for name, rating := range catalog {
		if rating > min {
			result = append(result, name)
		}
	}
	return result
}

func avgRating(catalog map[string]float64) float64 {
	if len(catalog) == 0 {
		return 0
	}

	totalR := 0.0
	quant := len(catalog)

	for _, rating := range catalog {
		totalR += rating
	}

	result := totalR / float64(quant)

	return result
}

func main() {
	catalog := map[string]float64{"Matrix": 9.2, "Inception": 8.8, "Star Wars": 9.0, "Protivostoyanie": 2.0}
	fmt.Println(catalog)

	newFilm := map[string]float64{"Hatiko": 9.0}
	newFilm1 := map[string]float64{"Matrix": 9.2}

	addFilm(newFilm, catalog)

	fmt.Println(catalog)

	addFilm(newFilm1, catalog)

	fmt.Println(catalog)

	fmt.Println(topFilms(catalog, 1.0))

	fmt.Println(avgRating(catalog))
}
