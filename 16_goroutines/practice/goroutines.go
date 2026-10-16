package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func checkEpisode(title string, delay int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(delay) * time.Microsecond)
	fmt.Printf("%s готов к публикации\n", title)
}

func main() {
	episodes := []string{"Episode1", "Desert", "Shitty shit", "Asshole world"}
	var wg sync.WaitGroup

	for _, episode := range episodes {
		wg.Add(1)
		go checkEpisode(episode, rand.Intn(1000), &wg)
	}
	wg.Wait()
	fmt.Println("все эпизоды проверены")
}
