package main

import "fmt"

const (
	price = 999.0
	tarif = "Premium"
)

func main() {
	var month int = 12

	var discount = 15.0

	total := (price - (price/100)*discount) * float64(month)

	fmt.Println("===== ЧЕК GoFlix =====")
	fmt.Printf("Тариф: %s \nМесяцев: %d \nБазовая цена: %.2f руб/мес \nСкидка: %.0f%% \nИтого: %.2f руб. ", tarif, month, price, discount, total)

}
