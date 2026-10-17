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

func transcodeVideo(name string, ch chan string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
	ch <- fmt.Sprintf("%s обработан заебись", name)
}

func main() {
	fmt.Println("1")
	episodes := []string{"Episode1", "Desert", "Shitty shit", "Asshole world"}
	var wg sync.WaitGroup

	for _, episode := range episodes {
		wg.Add(1)
		go checkEpisode(episode, rand.Intn(1000), &wg)
	}
	wg.Wait()
	fmt.Println("все эпизоды проверены")

	fmt.Println("2")
	videos := []string{"Movie", "Sex", "Shit", "Govno", "Humor", "Comedy"}

	channel := make(chan string)
	var wg2 sync.WaitGroup

	for _, video := range videos {
		wg2.Add(1)

		go func(v string) {
			defer wg2.Done()

			transcodeVideo(v, channel)
		}(video)
	}

	go func() {
		wg2.Wait()
		close(channel)
	}()

	for obrabotannoe_video := range channel {
		fmt.Println(obrabotannoe_video)
	}
}
