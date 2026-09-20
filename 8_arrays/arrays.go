package main

import "fmt"

func main() {
	//arrays
	var a [5]int                           // пять нулей
	b := [3]string{"Go", "Python", "Rust"} // сразу с данными
	c := [...]int{1, 2, 3}                 // компилятор сам посчитает

	fmt.Println(a, b, c)

	//slices
	s1 := []int{1, 2, 3}     // len=3, cap=3
	s2 := make([]int, 5)     // len=5, cap=5 -> [0 0 0 0 0]
	s3 := make([]int, 0, 10) // len=0, cap=10 -> []
	s4 := a[1:4]             // len=3, cap зависит от массива

	fmt.Println(s1, s2, s3, s4)

	original := []int{1, 2, 3, 4, 5}
	slice := original[1:4]
	slice[0] = 99
	fmt.Println(original) // [1 99 3 4 5] <- оригинал изменился!

	original = []int{1, 2, 3}
	clone := make([]int, len(original))
	copy(clone, original)
	clone[0] = 99
	fmt.Println(original)
	fmt.Println(clone)
}
