package main

import "fmt"

func main() {
	m1 := map[string]int{"Go": 1, "Python": 2} // с данными
	m2 := make(map[string]int)                 // пустая
	var m3 map[string]int                      // nil, опасно!

	fmt.Println(m1, m2, m3)

	m := make(map[string]float64)
	m["Дюна"] = 8.1
	m["Довод"] = 7.6
	m["Дюна"] = 8.3           // перезаписали
	fmt.Println(m["Дюна"])    // 8.3
	fmt.Println(m["Матрица"]) // 0

	rating, ok := m["Дюна"]
	fmt.Println(rating, ok) // 8.3 true

	rating, ok = m["Матрица"]
	fmt.Println(rating, ok) // 0 false

	if !ok {
		fmt.Println("Фильм не найден")
	}

	delete(m, "Довод")          // удалили
	delete(m, "Несуществующий") // ключа нет, ничего не произошло
	fmt.Println(m)
}
