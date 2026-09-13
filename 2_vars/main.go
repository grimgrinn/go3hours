package main

import "fmt"

func main() {
	var rating float64 = 8.6

	var title string = "string"

	var year int = 2014

	var matched bool = true

	fmt.Println(rating, title, year, matched)

	var rating2 = 8.5

	fmt.Println(rating2)

	rating3 := 8.1

	fmt.Println(rating3)

	rating3 = 8.2

	fmt.Println(rating3)

	var rating4 float64

	fmt.Println(rating4)

	const maxRating = 10.0

	fmt.Println(maxRating)

	const (
		maxRating1 = 11.0
		plan       = 1
	)

	width, height, fps := 640, 480, 60

	fmt.Println(width, height, fps)

	left, right := "left", "right"

	fmt.Println(left, right)

	left, right = right, left

	fmt.Println(left, right)

}
