package main

import (
	"fmt"
)

// 1
func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// 2
func removeDuplicates(s []string) []string {
	result := []string{}

	for _, str := range s {
		exists := false
		for _, r := range result {
			if r == str {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, str)
		}
	}
	return result
}

// 3
func filterByRating(ratings []float64, min float64) []float64 {
	result := make([]float64, 0, len(ratings))

	for _, rating := range ratings {
		if rating >= min {
			result = append(result, rating)
		}
	}
	return result
}

func main() {
	//1
	nums := []int{1, 2, 3, 4, 5}
	// nums := []int{}
	// var nums []int

	fmt.Println(sum(nums))

	//2
	s := []string{"apple", "banana", "apple", "orange", "banana"}
	s1 := []string{"a", "a", "a", "a"}
	s2 := []string{}

	fmt.Println(s, removeDuplicates(s))
	fmt.Println(s1, removeDuplicates(s1))
	fmt.Println(s2, removeDuplicates(s2))

	//3
	ratings := []float64{7.2, 5.1, 8.9, 6.0, 9.1, 4.3, 7.8}
	min := 7.8
	fmt.Println(ratings, min, filterByRating(ratings, min))

	ratings1 := []float64{1.0, 2.3, 3.4}
	min1 := 5.1
	fmt.Println(ratings1, min1, filterByRating(ratings1, min1))

	ratings2 := []float64{}
	min2 := 0.0
	fmt.Println(ratings2, min2, filterByRating(ratings2, min2))
}
