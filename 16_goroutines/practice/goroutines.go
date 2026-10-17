package main

import (
	"context"
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

func loadUserTaste(ch chan string) {
	sl := rand.Intn(25905)
	time.Sleep(time.Duration(sl * int(time.Microsecond)))
	ch <- fmt.Sprintf("Пользователь предпочитает %d говна", sl)
}

type Event struct {
	ID   string
	Type string
}

func worker(ctx context.Context, id int, jobs <-chan Event, results chan<- string, wg *sync.WaitGroup, mu *sync.Mutex, count *int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-jobs:
			if !ok {
				return
			}
			time.Sleep(time.Duration(rand.Intn(50)) * time.Microsecond)
			results <- fmt.Sprintf("Событие %s обработано: %s", event.ID, event.Type)
			mu.Lock()
			*count++
			mu.Unlock()
		}
	}
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

	fmt.Println("3")

	ch := make(chan string, 10)

	events := []string{"play", "pause", "seek", "stop"}

	for _, e := range events {
		ch <- e
	}

	close(ch)

	for e := range ch {
		fmt.Println(e)
	}

	fmt.Println("Очередь событий обработана")

	fmt.Println("4")

	userPreps := make(chan string)

	go loadUserTaste(userPreps)

	select {
	case res := <-userPreps:
		fmt.Println(res)
	case <-time.After(5009 * time.Microsecond):
		fmt.Println("Предпочтения пользователья пока недоступны")
	}

	fmt.Println("5")

	events1 := []Event{
		{ID: "play", Type: "control"},
		{ID: "pause", Type: "control"},
		{ID: "stop", Type: "control"},
		{ID: "like", Type: "reaction"},
		{ID: "fav", Type: "reaction"},
	}

	jobs := make(chan Event)
	results := make(chan string)

	var wg3 sync.WaitGroup
	var processedCount int
	var mu sync.Mutex

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg3.Add(1)
		go worker(ctx, i, jobs, results, &wg3, &mu, &processedCount)
	}

	go func() {
		for _, event := range events1 {
			jobs <- event
		}
		close(jobs)
	}()

	go func() {
		wg3.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Обработка прервана по таймауту, успели обработать:", processedCount)
	} else {
		fmt.Println("Успешно обработано событий:", processedCount)
	}

}
