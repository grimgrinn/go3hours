package main

import "fmt"

func main() {
	//for
	for ep := 1; ep <= 5; ep++ {
		fmt.Println(ep)
	}

	// while
	ep := 0
	for ep < 5 {
		fmt.Println(ep)
		ep++
	}

	//endless
	ep1 := 0
	for {
		fmt.Println(ep1)
		ep1++

		if ep1 > 5 {
			break
		}
	}

	// continue
	for i := 1; i <= 10; i++ {
		if i%3 != 0 {
			continue
		}
		fmt.Println(i)
	}

	//range
	genres := []string{"Боевик", "Комедия", "Драма", "Триллер", "Фантастика"}

	for idx, g := range genres {
		fmt.Println(idx, g)
	}

	// strings
	word := "Фильм"
	for pos, ch := range word {
		fmt.Printf("Позиция %d: '%c' (код %d)\n", pos, ch, ch)
	}

	// maps
	ratings := map[string]float64{
		"Начало":        8.8,
		"Интерстеллар":  8.7,
		"Темный рыцарь": 9.0,
	}

	for film, score := range ratings {
		fmt.Println(film, score)
	}

	//nested
	for row := 1; row <= 3; row++ {
		for seat := 1; seat <= 3; seat++ {
			fmt.Println(row, seat)
		}
	}

	//multifor
	seq := []int{10, 20, 30, 40, 50}

	for lo, hi := 0, len(seq)-1; lo < hi; lo, hi = lo+1, hi-1 {
		fmt.Println(lo, seq[lo], hi, seq[hi])
	}
}
