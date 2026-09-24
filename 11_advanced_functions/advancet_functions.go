package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Вариативные функции: базовый синтаксис ===")

	fmt.Printf(" Средний: %.1f\n", averageRating(8.0, 9.2, 7.5))
	fmt.Printf(" Одна оценка: %.1f\n", averageRating(9.0))
	fmt.Printf(" Без оценок: %.1f\n", averageRating())

	fmt.Println("\n=== Внутри это слайс ===")

	showVariadicInfo(7.0, 8.5, 9.2, 6.3)

	userRatings := []float64{8.2, 7.8, 9.1, 8.5}
	// averageRating(userRatings) // ОШИБКА
	avg := averageRating(userRatings...) // Правильно - разворачиваем слайс
	fmt.Printf("Из слайса: %.1f\n", avg)

	fmt.Println("\n=== Обычные + вариативные параметры ===")

	addToPlaylist("Mой плейлист", "Матрица", "Начало", "Дюна")
	addToPlaylist("Пустой плейлист")

	fmt.Println("\n===...any (любые типы) ===")
	debugPrint("Гладиатор", 2014, 8.6, true)

	fmt.Println("\n=== append - вариативная функция ===")

	films := []string{"A"}
	films = append(films, "B", "C")
	extra := []string{"D", "E"}
	films = append(films, extra...)
	fmt.Println(films)

	fmt.Println("\n=== Анонимная функция (IIFE) ===")

	func() {
		fmt.Println("Привет из анонимной функции")
	}()

	func(title string) {
		fmt.Println("Показываем:", title)
	}("Гладиатор")

	fmt.Println("\n=== Сохранение в переменную ===")

	formatMovie := func(title string, year int) string {
		return fmt.Sprintf("'%s' (%d)", title, year)
	}
	fmt.Println(formatMovie("Матрица", 1991))

	// можно переназначить
	formatMovie = func(title string, year int) string {
		return fmt.Sprintf("[%d] %s", year, strings.ToUpper(title))
	}
	fmt.Println(formatMovie("Довод", 2020))

	fmt.Println("\n=== Замыкания ===")

	// Функция захватывает переменные из окружения
	watchCount := 0
	watchMovie := func(title string) {
		watchCount++
		fmt.Printf("Просмотр #%d: %s\n", watchCount, title)
	}
	watchMovie("Гладиатор")
	watchMovie("Начало")
	fmt.Println("Всего:", watchCount) // 2

	fmt.Println("\n=== Генератор ID ===")

	newMovieID := makeIDGenerator("FILM")
	newUserID := makeIDGenerator("USER")
	fmt.Println(newMovieID()) // FILM-1
	fmt.Println(newMovieID()) // FILM-2
	fmt.Println(newUserID())  // USER-1
	fmt.Println(newUserID())  // USER-2
	fmt.Println(newUserID())  // USER-3

	fmt.Println("\n=== Callback (sort.Slice) ===")

	ratings := []float64{7.2, 9.1, 6.5, 8.8, 5.3}
	sort.Slice(ratings, func(i, j int) bool {
		return ratings[i] > ratings[j] // по убыванию
	})
	fmt.Println("Отсортированно:", ratings)

	fmt.Println("\n=== Функция высшего порядка ===")

	applyDisount := makeDiscountApplier(0.20)
	fmt.Printf("Со скидкой 20%%: %.2f\n", applyDisount(1000))

	titles := []string{"Матрица", "Начало", "Дюна"}
	fmt.Println("Капс:", transformStrings(titles, strings.ToUpper))

	benchmarkOperation("загрузка", func() {
		time.Sleep(50 * time.Millisecond)
	})
}

// makeDiscountApplier возвращает функцию, применяющую скидку.
func makeDiscountApplier(rate float64) func(float64) float64 {
	return func(price float64) float64 {
		return price * (1 - rate)
	}
}

// transformStrings приименяет функцию к каждой строке.
func transformStrings(items []string, fn func(string) string) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}

// benchmarkOperation замеряет время выполнения функции.
func benchmarkOperation(name string, op func()) {
	start := time.Now()
	op()
	fmt.Printf("%s: %v\n", name, time.Since(start))
}

// makeIDGenerator создаёт генератор уникальных ID c заданным префиксом
func makeIDGenerator(prefix string) func() string {
	counter := 0
	return func() string {
		counter++
		return fmt.Sprintf("%s-%d", prefix, counter)
	}
}

// addToPlaylist добавляет фильмы в плейлист.
func addToPlaylist(playlistName string, movies ...string) {
	if len(movies) == 0 {
		fmt.Printf(" '%s': пусто\n", playlistName)
		return
	}
	fmt.Printf(" '%s': %v\n", playlistName, movies)
}

// debugPrint выводит аргументы любых типов с их типами.
func debugPrint(args ...interface{}) {
	for i, arg := range args {
		fmt.Printf(" [%d] %v (%T)\n", i, arg, arg)
	}
}

// showVariadicInfo показывает тип, длину и содержимое вариативного параметра.
func showVariadicInfo(ratings ...float64) {
	fmt.Printf(" Тип: %T, Длина: %d, Содержимое: %v\n", ratings, len(ratings), ratings)
}

// averageRating считает средний рейтинг произвольного числа оценок.
func averageRating(ratings ...float64) float64 {
	if len(ratings) == 0 {
		return 0
	}
	sum := 0.0

	for _, r := range ratings {
		sum += r
	}

	return sum / float64(len(ratings))
}
