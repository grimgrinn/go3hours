package main

import "fmt"

func calculateDiscount(yearsSubscribed int) int {
	if yearsSubscribed >= 10 {
		return 20
	} else if yearsSubscribed >= 5 {
		return 10
	} else if yearsSubscribed >= 1 {
		return 5
	}
	return 0
}

func categorizeRating(r float64) string {
	if r >= 8.5 {
		return "The best"
	} else if r >= 7.0 {
		return "Good"
	} else if r >= 5.0 {
		return "Okay"
	} else {
		return "Bad"
	}
}

func main() {

	if d := calculateDiscount(12); d > 0 {
		fmt.Println(d)
	} else {
		fmt.Println("no discount")
	}

	ticketsLeft := -2

	if ticketsLeft > 0 {
		fmt.Println("билеты есть")

	} else if ticketsLeft == 0 {
		fmt.Println("билетов нет")
	} else {
		fmt.Println("Система сломаласб")
	}

	// && и
	userAge := 22
	hasPremium := true
	if userAge > 18 && hasPremium {
		fmt.Println("good")
	}
	// || или
	userAge = 17
	hasPremium = true
	if userAge > 18 || hasPremium {
		fmt.Println("good")
	}
	// ! не
	fmt.Println(!true)

	isLoggedIn := true
	isVerified := true
	hasTicket := true

	if isLoggedIn && isVerified && hasTicket {
		fmt.Println("good")
	}

	films := []float64{9.2, 7.5, 5.8, 3.2, 8.0}

	for _, r := range films {
		fmt.Println(categorizeRating(r))
	}
}
