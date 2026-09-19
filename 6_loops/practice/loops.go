package main

import "fmt"

func main() {

	var films = []string{"Дюна", "Матрица", "Начало", "Интерстеллар"}
	var seanses = []string{"10:00", "13:00", "16:00", "19:00"}
	var halls = []int{1, 2, 3}

	for h := range halls {
		fmt.Printf("Зал %d: \n", h)
		for s := range seanses {
			fmt.Printf("%s -  %s \n", seanses[s], films[s])

		}
	}

	// for h := 0; h < len(halls); h++ {
	// 	fmt.Println(halls[h])
	//
	// for s := 0; s < len(seanses); s++ {
	// 	fmt.Println(seanses[s])
	// }

	// for f := 0; f < len(films); f++ {
	// 	fmt.Println(films[f])
	// }
	count := 0
	for le := 0; le < len(films); le++ {
		for f := range films[le] {
			fmt.Printf("%d", f)
			count++
		}
		fmt.Printf("%s - %d runes\n", films[le], count)
	}

}
