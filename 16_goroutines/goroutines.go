package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ViewCounter - потокобезопасный счетчик просмотров
type ViewCounter struct {
	mu    sync.Mutex
	count int
}

func (vc *ViewCounter) Increment() {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	vc.count++
}
func (vc *ViewCounter) Value() int {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	return vc.count
}

// worker - воркер, обрабатывающий задачи из канала jobs.
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for film := range jobs {
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		results <- fmt.Sprintf("Воркер %d обработал: %s", id, film)
	}
}

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

	fmt.Println("=======Select=========")

	recomendations := make(chan string)
	go func() {
		time.Sleep(50 * time.Millisecond)
		recomendations <- "Интерстеллар"
	}()

	select {
	case movie := <-recomendations:
		fmt.Println(" Рекомендации:", movie)
	case <-time.After(200 * time.Millisecond):
		fmt.Println(" Таймаут! Сервис рекомендаций не ответил")
	}

	// Неблокирующий select с default
	ch3 := make(chan string)
	select {
	case msg := <-ch3:
		fmt.Println(msg)
	default:
		fmt.Println(" Нет данных - идем дальше")
	}

	fmt.Println("========Sync.Mutex=======")

	counter := &ViewCounter{}
	var wg2 sync.WaitGroup

	// Запускаем 100 горутин, каждая увеличивает счетчик
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			counter.Increment()
		}()
	}
	wg2.Wait()
	fmt.Printf(" Просмотров (100 горутин): %d\n", counter.Value())

	fmt.Println("=======Worker Pool=========")

	filmList := []string{"Дюна", "Иатрица", "Довод", "Начало", "Аватар"}

	jobs := make(chan string)
	results := make(chan string)

	var wg3 sync.WaitGroup

	// Запускаем три воркера
	for w := 1; w <= 3; w++ {
		wg3.Add(1)
		go worker(w, jobs, results, &wg3)
	}
	// Отправляем задачи
	go func() {
		for _, film := range filmList {
			jobs <- film
		}
		close(jobs) // сигнал воркерам, что задач больше нет
	}()
	// Закрываем results только после завершения всех воркеров
	go func() {
		wg3.Wait()
		close(results)
	}()
	// Собираем результаты
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("====Context с таймаутом====")

	// Создаем контекст с таймаутом 150 мс
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel() // освобождаем ресурсы после завершения

	resultCh := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond) // "долгая" операция
		resultCh <- "Рекомендации готовы!"
	}()

	select {
	case res := <-resultCh:
		fmt.Println(res)
	case <-ctx.Done():
		fmt.Println("Отмена:", ctx.Err())
	}

	fmt.Println("\n=== Context с отменой ===")

	ctx2, cancel2 := context.WithCancel(context.Background())
	var wg4 sync.WaitGroup
	wg.Add(1)

	// Горутина должна реагировать на отмену через ctx.Done()
	go func(ctx context.Context) {
		defer wg4.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Стриминг остановлен")
				return
			default:
				fmt.Println("Стрими фильм...")
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(ctx2)
	time.Sleep(120 * time.Millisecond)
	cancel2() // отменяем - горутина завершится
	wg.Wait()
}
