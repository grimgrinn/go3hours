package main

import "fmt"

var totalFilms = 0

const maxCatalogSize = 10000

func addFilm(number int) int {
	totalFilms += number
	return totalFilms
}

func changeNumber(number int) (int, int) {
	number -= 10
	return number, number - 20
}

func ticketPrice(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(totalFilms)

	outerVal := "value1"

	{
		innerVal := "value2"
		fmt.Println(outerVal)
		fmt.Println(innerVal)

	}

	fmt.Println(outerVal)
	//fmt.Println(innerVal) // ошибка

	price := 100
	fmt.Println(price)
	{
		price := 200
		fmt.Println(price)
	}

	fmt.Println(price)

	addFilm(100)
	fmt.Println(addFilm(100))

	number := 500
	fmt.Println(changeNumber(number))
	fmt.Println(number)

	a, b := changeNumber(number)

	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)

	var calc func(int, int) int

	calc = ticketPrice

	fmt.Println(calc(1, 2))

}
