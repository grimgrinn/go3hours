package main

import "fmt"

const movie_length_seconds = 8520

func main() {
	total_minutes := movie_length_seconds / 60
	rem_seconds := movie_length_seconds % 60
	total_hours := total_minutes / 60
	rem_minutes := total_minutes % 60

	fmt.Printf("%d seconds is %d hours %d minutes %d seconds", movie_length_seconds, total_hours, rem_minutes, rem_seconds)
}
