package main

import "fmt"

func checkAccess(age int, isPremium, hasBan bool, rating int) (bool, string) {
	if hasBan {
		return false, "Аккаунт забанен"
	}

	if age < rating {
		return false, "Не проходишь по возрасту"
	}

	if isPremium {
		return true, "Приятного просмотра"
	}

	return true, "Доступно с рекламой"
}

func main() {
	age := 16
	isPremium := false
	hasBan := false
	rating := 18

	fmt.Println(checkAccess(age, isPremium, hasBan, rating))

	age = 26
	isPremium = false
	hasBan = false
	rating = 18

	fmt.Println(checkAccess(age, isPremium, hasBan, rating))

	age = 36
	isPremium = true
	hasBan = false
	rating = 18

	fmt.Println(checkAccess(age, isPremium, hasBan, rating))

	age = 36
	isPremium = true
	hasBan = true
	rating = 18

	fmt.Println(checkAccess(age, isPremium, hasBan, rating))

	if d, s := checkAccess(18, false, false, 6); !d {
		fmt.Println("doen not!", s)
	} else {
		fmt.Println("it does!", s)
	}
}
