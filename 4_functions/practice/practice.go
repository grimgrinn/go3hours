package main

import "fmt"

func ticketPrice(price int, evening, holiday bool) int {
	if evening {
		price += 50
	}

	if holiday {
		price += 100
	}

	return price
}

func formatReceipt(title string, quantity, price int) (string, int) {
	total := price * quantity
	return title, total
}

func isAllowed(age, pr int) bool {
	if age < pr {
		return false
	}

	return true
}

func main() {
	fmt.Println("Ticket price:  price 100, evening, sunday\n", ticketPrice(100, true, true))

	name, total := formatReceipt("Odisseyd", 3, 399)
	fmt.Printf("formatReceipt: Odisseya, 3, 399: %s, %d\n", name, total)

	fmt.Printf("isAllowed: age 16 pr 18, age 16 pr 13: %v, %v", isAllowed(16, 18), isAllowed(16, 13))

}
