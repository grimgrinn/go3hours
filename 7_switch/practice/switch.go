package main

import "fmt"

func identifyType(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("%v -> int\n", val)
	case string:
		fmt.Printf("%v -> string\n", val)
	case float64:
		fmt.Printf("%v -> float64\n", val)
	case bool:
		fmt.Printf("%v -> bool\n", val)
	default:
		fmt.Printf("%v -> %T\n", val, val)
	}
}

func main() {
	commands := []string{"play", "pause", "stop", "next", "prev", "volume_up", "volume_down", "info", "like", "quit", "share"}

	for _, command := range commands {
		switch command {
		case "play":
			fmt.Println("play")
		case "pause":
			fmt.Println("pause")
		case "stop":
			fmt.Println("stop")
		case "next":
			fmt.Println("next")
		case "prev":
			fmt.Println("prev")
		case "volume_up":
			fmt.Println("volume_up")
		case "volume_down":
			fmt.Println("volume_down")
		case "info":
			fmt.Println("info")
		case "like":
			fmt.Println("like")
		case "quit":
			fmt.Println("quit")
		case "share":
			fmt.Println("share")
		}
	}

	identifyType(43)
	identifyType("Матрица")
	identifyType(9.1)
	identifyType(true)
}
