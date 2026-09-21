package main

import "fmt"

func countChars(s string) map[rune]int {
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}
	return counts
}

func hasDuplicates(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

func main() {
	s := "hello world"

	for ss, s1 := range countChars(s) {
		fmt.Printf("%c: %d\n", ss, s1)
	}

	nums := []int{1, 2, 3, 1}
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 43, 6, 34, 1}
	nums3 := []int{2, 43, 6, 34, 1, 43}

	fmt.Println(hasDuplicates(nums))
	fmt.Println(hasDuplicates(nums1))
	fmt.Println(hasDuplicates(nums2))
	fmt.Println(hasDuplicates(nums3))

}
