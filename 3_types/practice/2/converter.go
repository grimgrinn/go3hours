package main

import "fmt"

const size_bytes = 4_831_838_208
const bytes_in_gb = 1_073_741_824

func main() {
	var gbs float64

	gbs = float64(size_bytes) / float64(bytes_in_gb)

	fmt.Printf("Размер в гигабайтах = %.2f", gbs)
}
