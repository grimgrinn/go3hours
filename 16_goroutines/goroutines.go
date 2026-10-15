package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Пример: запуск нескольких горутин
	fmt.Println("=== Горутины ===")

	go fmt.Println("Привет из горутины!")

	go func() {
		fmt.Println("Привет из анонимной горутины!")
	}()

	// Sleep лишь для демонстрации.
	time.Sleep(100 * time.Microsecond)

	fmt.Println("========WaitGroup==========")
	var wg sync.WaitGroup
	movies := []string{"Интерстеллар", "Дюна", "Начало"}

	for _, movie := range movies {
		wg.Add(1) // увеличиваем счетчик перед запуском
		go func(title string) {
			defer wg.Done() // уменьшаем счетчик при заершении
			duration := time.Duration(rand.Intn(100)) * time.Microsecond

			time.Sleep(duration)
			fmt.Printf(" Загружаем постер: %s (%v)\n", title, duration)
		}(movie)
	}
	wg.Wait() // ждём, пока все горутины не закончат
	fmt.Println(" Все постеры загружены!")

	fmt.Println("========Unbuffered Channels==========")

	ch := make(chan string) // Создаём небуферизованный канал строк

	//Отправитель
	go func() {
		ch <- "Mатрица загружена!" // Блокируется, пока кто-то не прочтет
	}()

	// Получатель
	msg := <-ch // ждет, пока придет значение
	fmt.Println(msg)

	// Канад с числовым результатом
	ratingCh := make(chan float64)
	go func() {
		ratings := []float64{8.0, 9.2, 7.5, 8.8}
		sum := 0.0
		for _, r := range ratings {
			sum += r
		}
		ratingCh <- sum / float64(len(ratings))
	}()
	fmt.Printf("Средний рейтинг: %.1f\n", <-ratingCh)

	fmt.Println("========Buffered Channels=======")

	notifications := make(chan string, 3)

	notifications <- "Hoвый фильм: Дюна 2"
	notifications <- "Скидка 50%"
	notifications <- "Новый сезон: Мандалорец"

	// notifications <- "Еще одно" // заблокируется - буфер полон

	fmt.Println(len(notifications), "уведомления в очереди")
	fmt.Println(<-notifications)
	fmt.Println(<-notifications)
	fmt.Println(<-notifications)

	fmt.Println("=======Range и close=========")

	movieCh := make(chan string)
	go func() {
		for _, m := range []string{"Матрица", "Начало", "Довод"} {
			movieCh <- m
		}
		close(movieCh) // закрываем канал - сигнал для range
	}()

	for movie := range movieCh {
		fmt.Println(" Получен:", movie)
	}

	// Демонстрация comma-ok: чтение из закрытого канала
	ch2 := make(chan int)
	close(ch2)
	val, ok := <-ch2 // val = 0, ok = false - канал закрыт и пуст
	fmt.Println("Значение:", val, "Открыт?:", ok)

	fmt.Println("=======Сигнальный канал=======")

	done := make(chan struct{})

	go func() {
		fmt.Println("Горутина работает...")
		time.Sleep(100 * time.Millisecond)
		done <- struct{}{}
	}()

	// Ждем сигнал
	<-done
	fmt.Println("main: получили сигнал завершения")
}
