package main

import (
	"fmt"
	"time"
)

func handleAction(action string) {
	switch action {
	case "play":
		fmt.Println(">")
	case "pause":
		fmt.Println("||")
	default:
		fmt.Println("No such actdion")
	}
}

func main() {
	genre := 3

	switch genre {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println(5)
	}

	contentRating := "PG-13"

	switch contentRating {
	case "G", "PG":
		fmt.Println(1)
	case "PG-13":
		fmt.Println(2)
	}

	switch h := time.Now().Hour(); {
	case h < 12:
		fmt.Println("Утро")
	case h < 13:
		fmt.Println("День")
	default:
		fmt.Println("Вечер/Ночь")
	}

	score := 7.3

	switch {
	case score >= 9.0:
		fmt.Println(1)
	case score >= 7.5:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	tier := 1
	switch tier {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	default:
		fmt.Println("3")
	}

	actions := []string{"play", "pause", "like", "quit", "share"}

	for _, action := range actions {
		handleAction(action)
	}
}
