package main

import (
	"fmt"
	"math"
)

// Content - интерфейс "контента": всё, что имеет название, длительность, и описание.
type Content interface {
	Title() string
	Duration() int
	Description() string
}

// Playable - умеет воспроизводиться.
type Playable interface {
	Play() string
	Pause() string
}

// Series реализует Content и Playable.
type Series struct {
	Name     string
	Seasons  int
	Episodes int
	AvgLen   int
}

func (s Series) Title() string { return s.Name }
func (s Series) Duration() int { return s.Episodes * s.AvgLen }
func (s Series) Description() string {
	return fmt.Sprintf("%s, %d сезонов", s.Name, s.Seasons)
}
func (s Series) Play() string  { return fmt.Sprintf("> '%s' S01E01", s.Name) }
func (s Series) Pause() string { return fmt.Sprintf("|| '%s'", s.Name) }

// Podcast реализует только Content (нет методов Play/Pause).
type Podcast struct {
	Name    string
	Host    string
	Minutes int
}

func (p Podcast) Title() string { return p.Name }
func (p Podcast) Duration() int { return p.Minutes }
func (p Podcast) Description() string {
	return fmt.Sprintf("Подкаст '%s'  от %s", p.Name, p.Host)
}

type Movie struct {
	Name    string
	Year    int
	Rating  float64
	Genre   string
	Minutes int
}

func (m Movie) Title() string { return m.Name }
func (m Movie) Duration() int { return m.Minutes }
func (m Movie) Description() string {
	return fmt.Sprintf("Фильм %s %d мин", m.Name, m.Minutes)
}
func (m Movie) Play() string  { return fmt.Sprintf("> Play %s\n", m.Name) }
func (m Movie) Pause() string { return fmt.Sprintf("|| Pause  %s\n", m.Name) }

func printContentInfo(c Content) {
	fmt.Printf("%s - %d мин\n", c.Description(), c.Duration())
}

// Shape - геометрическая фигура.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle - приямоугольник.
type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Circle - круг.
type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

func classifyContent(v interface{}) {
	switch val := v.(type) {
	case Movie:
		fmt.Printf(" Фильм: %s (%d)\n", val.Name, val.Year)
	case string:
		fmt.Printf(" Строка: %s\n", val)
	case int:
		fmt.Printf(" Число: %d\n", val)
	case float64:
		fmt.Printf(" Float: %.1f\n", val)
	case bool:
		fmt.Printf(" Bool: %v\n", val)
	default:
		fmt.Printf(" Другое: %T\n", val)
	}
}

// MediaPlayer = Content + Playable
type MediaPlayer interface {
	Content
	Playable
}

// GoFlixUser реализует fmt.Stringer
type GoFlixUser struct {
	Name         string
	Plan         string
	MovieWatched int
}

func (u GoFlixUser) String() string {
	return fmt.Sprintf("{%s, %s, %d просмотров}", u.Name, u.Plan, u.MovieWatched)
}

// SubscriptionError - кастомный тип ошибки.
type SubscriptionError struct {
	Status string
	Code   int
}

// Реализация метода Error() - достаточно для удовлетворения интерфейса error.
func (e *SubscriptionError) Error() string {
	return fmt.Sprintf("подписка: '%s' (код %d)", e.Status, e.Code)
}

// Функция проверяет статус подписки и возвращает error.
func checkSubscription(status string) error {
	if status == "expired" {
		return &SubscriptionError{Status: "expired", Code: 402}
	}
	if status == "banned" {
		return &SubscriptionError{Status: "banned", Code: 403}
	}

	return nil
}

func main() {
	var item Content
	// Интерфейсная переменная может содержать фильм, сериал или подкаст
	item = Movie{Name: "Интерстеллар", Year: 2014, Minutes: 169, Genre: "Фантастика"}
	printContentInfo(item)
	item = Series{Name: "Во все тяжкие", Seasons: 5, Episodes: 62, AvgLen: 47}
	printContentInfo(item)

	// Полиморфизм: разные типы в одном слайсе
	catalog := []Content{
		Movie{Name: "Дюна", Year: 2021, Minutes: 155, Genre: "Фантастика"},
		Series{Name: "Чернобыль", Seasons: 1, Episodes: 5, AvgLen: 65},
		Podcast{Name: "GoFliz Talks", Host: "Анна", Minutes: 45},
	}
	total := 0
	for _, c := range catalog {
		fmt.Printf("%s - %d мин\n", c.Title(), c.Duration())
		total += c.Duration()
	}
	fmt.Printf(" Итого: %d мин\n", total)

	shapes := []Shape{
		Rectangle{Width: 16, Height: 9},
		Circle{Radius: 5},
	}

	for _, s := range shapes {
		fmt.Printf("Площадь: %.2f, Периметр: %.2f\n", s.Area(), s.Perimeter())
	}

	var anything interface{} // то же самое, что var anything any
	anything = "строка"
	fmt.Printf(" %v (%T)\n", anything, anything)
	anything = 42
	fmt.Printf(" %v (%T)\n", anything, anything)

	var i interface{} = "Привет"

	// Небезопасный вариант - паника, если тип неверный
	str := i.(string) // ok, i содержит строку
	fmt.Println(str)

	// Безопасный вариант - возвращает ok
	str2, ok := i.(string)
	fmt.Println(str2, ok) // ok = true

	num, ok := i.(int)
	fmt.Println(num, ok) // 0 false - i не int

	values := []interface{}{
		Movie{Name: "Матрица", Year: 1999, Minutes: 136, Genre: "Фантастика"},
		"строка", 42, 8.7, true,
	}

	for _, v := range values {
		classifyContent(v)
	}

	var player MediaPlayer = Movie{
		Name: "Начало", Year: 2010, Minutes: 148, Genre: "Фантастика",
	}
	fmt.Println(player.Description()) // доступно из Content
	fmt.Println(player.Play())        // доступно из Playable
	// Podcast не реализует Playable, поэтому нельзя присвоить его переменной MediaPlayer

	user := GoFlixUser{
		Name:         "Алексей",
		Plan:         "Premiun",
		MovieWatched: 45,
	}
	fmt.Println(user) // вызывает String() и печатает красиво

	err := checkSubscription("expired")
	if err != nil {
		fmt.Println("Ошибка", err)
		if subErr, ok := err.(*SubscriptionError); ok {
			// Type assertion для извлечения конкретного типа ошибки
			fmt.Printf("Статус: %s, Код: %d\n", subErr.Status, subErr.Code)
		}
	}

	// Интерфейс == nil только если и тип, и значение nil
	var c Content = nil
	fmt.Println("c == nil:", c == nil) // true

	// Опасность: присваиваем nil-указатель конкретного типа
	var movie2 *Movie = nil
	c = movie2
	fmt.Println("c == nil после *Movie(nil):", c == nil) // false!
	// c содержит (тип: *Movie, значеие: nil). Проверка на nil не сработает

	// При вызове метода возникает паника, потому что значение nil
	// fmt.Println(c.Title()) // panic: runtime error: invalid memory address
}
