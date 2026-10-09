package main

import (
	"fmt"
	"os"
	"time"
)

func demoBasicDefer() {
	fmt.Println("Начало")
	defer fmt.Println("[defer] Конец demoBasicDefer")
	fmt.Println("Продолжаем...")
}

func demoDeferArguments() {
	filmID := 100
	defer fmt.Println(" [defer аргумент] filmD =", filmID)

	filmID = 200
	fmt.Println("  сейчас filmID =", filmID)
}

func demoDeferClosure() {
	viewCount := 500
	defer func() {
		fmt.Println(" [defer замыкание] viewCount = ", viewCount)
	}()

	viewCount = 999
	fmt.Println(" сейчас viewCount = ", viewCount)
}

func demoLIFO() {
	defer fmt.Println(" [1] первый defer - последний")
	defer fmt.Println(" [2] второй defer")
	defer fmt.Println(" [3] третий defer - первый")

	fmt.Println("  обычный код")
}

func trackTime(start time.Time, name string) {
	fmt.Printf("  %s: %v\n", name, time.Since(start).Round(time.Microsecond))
}

func loadMovieCatalog() {
	defer trackTime(time.Now(), "Загрузка каталога")
	time.Sleep(80 * time.Millisecond)
	fmt.Println("  каталог загружен")
}

func writeToTempFile() {
	file, err := os.CreateTemp("", "goflix_catalog_*.txt")
	if err != nil {
		fmt.Println("  ошибка создания файла:", err)
		return
	}

	defer os.Remove(file.Name())
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("  ошибка закрытия файла:", err)
			return
		}

		fmt.Println(" файл закрыт")
	}()

	fmt.Println(" файл создан")

	if _, err := file.WriteString("GoFlix catalog data"); err != nil {
		fmt.Println("  ошибка записи:", err)
		return
	}

	fmt.Println("  данные записаны")
}

func demoPanicFlow() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(" [recover] паника перехвачена:", r)
		}
	}()

	defer fmt.Println("  [defer demoPanicFlow] освобождаем контекст запроса")

	startEpisodePlayback()

	fmt.Println("  жто строка не выполнитмя")
}

func startEpisodePlayback() {
	defer fmt.Println(" [defer playback] закрываем буфер")

	fmt.Println("  стартуем восароизведение")
	panic("Битый файл эпизода")
}

func safeExecute(name string, fn func()) (result string) {
	defer func() {
		if r := recover(); r != nil {
			result = fmt.Sprintf("%s: паника перехвачена: %v", name, r)
		}
	}()

	fn()

	return fmt.Sprintf("%s: успешно", name)
}

func handleRequest(movieTitle string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("внутренняя ошибка обработчика: %v", r)
		}
	}()

	if movieTitle == "" {
		panic("movieTitle пустой после валидации")
	}

	fmt.Printf("  запрос обработан: %q\n", movieTitle)
	return nil
}

func averageWatchTime(totalMinutes, viewers int) (int, error) {
	if viewers == 0 {
		return 0, fmt.Errorf("Количество зрителей равно нулю")
	}

	return totalMinutes / viewers, nil
}

func demoBadDeferInLoop(files []string) {
	for _, name := range files {
		fmt.Println(" открыли", name)
		defer fmt.Println(" закрыли", name)
	}

	fmt.Println("  конец функции demoBadDeferInLoop")
}

func processDataFile(name string) {
	fmt.Println(" открыли", name)
	defer fmt.Println(" закрыли", name)

	fmt.Println(" обработали", name)
}

func main() {
	fmt.Println("=== defer: базовай пример ===")
	demoBasicDefer()

	fmt.Println("=== Порядок defer: LIFO === ")
	demoLIFO()

	fmt.Println("=== Аргументы defer вычисляются сразу ===")
	demoDeferArguments()

	fmt.Println("=== defer c замыканием ===")
	demoDeferClosure()

	fmt.Println("=== Паттерн: открыл ресурс -> defer Close() ===")
	writeToTempFile()

	fmt.Println("=== defer для измерения времени ===")
	loadMovieCatalog()

	fmt.Println("=== panic и defer ===")
	demoPanicFlow()

	fmt.Println("=== recover как защитная обочка ===")

	result := safeExecute("request-42", func() {

		fmt.Println("  Загружаем рекомендации...")
		panic("Кэш рекомендаций поврежден")
	})

	fmt.Println("Результат:", result)
	fmt.Println("Программа продолжает работу")

	fmt.Println("=== recover на границе обработчика ===")

	for _, title := range []string{"Матрица", ""} {
		if err := handleRequest(title); err != nil {
			fmt.Println("  ошибка:", err)
		}
	}
	fmt.Println("  сервер продолжает принимать запросы")

	fmt.Println("=== panic vs error ===")

	avg, err := averageWatchTime(120, 0)

	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Среднее время просмотра:", avg)
	}

	fmt.Println("=== defer в цикле: ловушка ===")
	files := []string{"catalog.json", "users.json", "ratings.json"}
	demoBadDeferInLoop(files)

	fmt.Println("=== defer в цикле: правильный вариант ===")
	for _, name := range files {
		processDataFile(name)
	}
}
