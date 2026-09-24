package main

import "fmt"

func main() {
	fmt.Println("=== & и * ===")
	movieID := 42

	ptr := &movieID // ptr хранит АДРЕС movieID

	fmt.Println("movieID =", movieID)
	fmt.Println("ptr =", ptr)        // адрес в памяти
	fmt.Println("*ptr =", *ptr)      // 42 - значение по адресу
	fmt.Printf("Тип ptr: %T\n", ptr) // *int

	fmt.Println("\n=== Изменение через указатель ===")

	viewCount := 1000
	viewPtr := &viewCount
	*viewPtr = 5000
	fmt.Println("viewCount:", viewCount) // 5000

	fmt.Println("\n=== Указатели в функциях === ")

	// без указателя - копия, изменения не видны
	rating := 8.0
	tryChangeRating(rating)
	fmt.Println("После tryChange:", rating) // 8.0

	// с указателем - меняем ориинал
	changeRating(&rating, 9.5)
	fmt.Println("После change:", rating) // 9.5

	fmt.Println("\n=== Указатели на структуры ===")

	//Go автоматически разыменовывает: ptr.Field вместо (*ptr).Field
	film := Movie{Title: "Довод", Rating: 7.8, Views: 50000}
	addView(&film)
	addView(&film)
	fmt.Println(film) // Views = 50002

	filmPtr := &film
	fmt.Println(filmPtr.Title) // автоматическое разыменование

	fmt.Println("\n===  new() ===")

	numPtr := new(int) // *int, *numptr == 0
	*numPtr = 999
	fmt.Println(*numPtr)

	fmt.Println("\n=== nil-указатель ===")

	var nilPtr *int
	fmt.Println("nilPtr == nil:", nilPtr == nil)
	// _ = *nilPtr // panic!
	// if nilPtr != nil {
	// 	fmt.Println(*nilPtr)
	// }

	fmt.Println("\n=== Опциональные значения ===")

	// nil = "значение не задано" (отличие от нулевого значения)
	var userRating *float64
	fmt.Println("Не задан:", userRating) // nil

	score := 8.5
	userRating = &score
	printUserRating("Алексей", userRating)
	printUserRating("Мария", nil)

	fmt.Println("\n=== Возврат указателя === ")

	// в Go это безопасно (GC не удалит)
	newFilm := createMovie("Оппенгеймер", 8.5)
	fmt.Println(*newFilm)

	fmt.Println("\n=== Ccылочные типы ===")

	// слайсы, мапы, каналы - уже ссылочные
	films := []string{"A", "B", "C"}
	modifySlice(films)
	fmt.Println(films) // [ИЗМЕНЕН В С]

	scores := map[string]float64{"Дюна": 8.1}
	modifyMap(scores)
	fmt.Println(scores) // + Гладиатор

}

// tryChangeRating пытается изменить рейтинг - не получится (копия).
func tryChangeRating(r float64) {
	r = 10.0
}

// changeRating изменяет рейтинг через указатель - работает!
func changeRating(r *float64, newRating float64) {
	*r = newRating
}

type Movie struct {
	Title  string
	Rating float64
	Views  int
}

func addView(m *Movie) {
	m.Views++
}

// printUserRating выводит оценку пользователя ( или "не оценил").
func printUserRating(name string, rating *float64) {
	if rating != nil {
		fmt.Printf(" %s: %.1f\n", name, *rating)
	} else {
		fmt.Printf(" %s: не оценил\n", name)
	}
}

// createMovie создает фильм и возращает указателью
func createMovie(title string, rating float64) *Movie {
	m := Movie{Title: title, Rating: rating}
	return &m
}

// modifySlice изменяет первый элемент слайса.
func modifySlice(s []string) {
	if len(s) > 0 {
		s[0] = "ИЗМЕНЕН"
	}
}

// modifyMap добавляет элемент в мапу.
func modifyMap(m map[string]float64) {
	m["Гладиатор"] = 8.6
}
