package main

import "fmt"

func main() {
	var zeroInt int
	views := 1_500_000

	fmt.Println(zeroInt)
	fmt.Println(views)

	var zeroFloat float64
	views2 := 85.5

	fmt.Println(zeroFloat)
	fmt.Println(views2)

	var zeoString string
	movieTitle := "title"

	fmt.Println(zeoString)
	fmt.Println(movieTitle)

	sym := `sfdsfsdfsdfdsf
	sdfsdfsdfsdf
	sdfsdfsdf`

	fmt.Println(sym)

	var letterA rune = 'A'

	fmt.Println(letterA)

	fmt.Printf("%c\n", letterA)

	var zeroBool bool
	isPremium := true
	hasSubtitle := false

	fmt.Println(zeroBool)
	fmt.Println(isPremium)
	fmt.Println(hasSubtitle)

	episodes := 8
	avgLength := 52.5

	totalMinutes := float64(episodes) * avgLength

	fmt.Println(totalMinutes)
}
