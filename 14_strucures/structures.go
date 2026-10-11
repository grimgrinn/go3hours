package main

import (
	"fmt"
	"time"
)

// Film - описание фильма в GoFlix.
type Film struct {
	Title    string
	Year     int
	Rating   float64
	Genre    string
	Duration int
}

// Subscriber - подписчик GoFlix.
type Subscriber struct {
	Name  string
	Email string
	Age   int
	Plan  string
}

// Coordinates - пример структуры с полями одного типа.
type Coordinates struct {
	Lat, Lng float64
}

// FilmlCard - карточка фильма.
type FilmCard struct {
	Title     string
	Year      int
	Rating    float64
	Watched   bool
	ViewCount int
}

// Метод с получателем (ресивером) по значению: не изменяет исходную структуру.
func (f FilmCard) info() string {
	return fmt.Sprintf("'%s' (%d), рейтинг: %.1f", f.Title, f.Year, f.Rating)
}

// Метод с получателем по указателю: может менять структуру.
func (f *FilmCard) Watch() {
	f.Watched = true
	f.ViewCount++
}

// AdminUser - администратор сервиса. Содержит все поля Subscriber.
type AdminUser struct {
	Subscriber  // встраиваем Subscriber; доступ к Name, Email, Plan напрямую
	AccessLevel string
	Department  string
}

// CatalogItem - элемент каталога с тегами для JSON и базы данных.
type CatalogItem struct {
	ID        int     `json:"id" db:"item_id"`
	Title     string  `json:"title" db:"title"`
	Price     float64 `json:"price" db:"price"`
	Available bool    `json:"availalbe" db:"is_available"`
}

// StreamSession - сессия просмотра.
type StreamSession struct {
	UserName  string
	MovieName string
	StartedAt time.Time
	IsActive  bool
}

// NewStreamSession - конструктор StreamSession.
func NewStreamSesseion(user, movie string) *StreamSession {
	return &StreamSession{
		UserName:  user,
		MovieName: movie,
		StartedAt: time.Now(),
		IsActive:  true,
	}
}

func main() {

	var film Film
	film2 := Film{Title: "Интерстеллар", Year: 2014, Rating: 8.6, Genre: "Фантастика", Duration: 169}
	film3 := Film{Title: "Дюна", Year: 2021, Rating: 8.1, Genre: "Фантастика", Duration: 155}
	film4 := Film{Title: "Довод", Year: 2020}
	film5 := new(Film)
	fmt.Println(film, film2, film3, film4, film5)

	sub := Subscriber{Name: "Алексей", Email: "alex.hoover@goflix.ru", Age: 28, Plan: "Premium"}
	fmt.Println(sub.Name, sub.Plan) // читаем
	sub.Plan = "Standart"           // записываем

	filmPtr := &film2         // Берем адрес существующей структуры
	filmPtr.Rating = 8.7      // Go сам разыменовывает, * не нужен
	fmt.Println(film2.Rating) // 8.7 - оригинальная структура изменилась

	admin := AdminUser{
		Subscriber:  Subscriber{Name: "Елена", Email: "e@goflix.tu", Plan: "Premiun"},
		AccessLevel: "superadmin",
		Department:  "Billing",
	}

	fmt.Println(admin.Name)             // "Елена" - поле из Subscriber
	fmt.Println(admin.Subscriber.Email) // можно и так

	config := struct {
		Host string
		Port int
	}{
		Host: "api.goflix.ru",
		Port: 443,
	}
	fmt.Printf("%s:%d\n", config.Host, config.Port)

	// Табличные тесты (table-driven tests)
	tests := []struct {
		Input    string
		Expected int
	}{
		{Input: "A", Expected: 7},
		{Input: "B", Expected: 9},
	}
	fmt.Println(tests)

	genres := map[string]struct{}{
		"Фантастика": {},
		"Драма":      {},
		"Комедия":    {},
	}

	// Добавляем новый жанр
	genres["Боевик"] = struct{}{}

	// Проверяем наличие
	if _, exists := genres["Драма"]; exists {
		fmt.Println("Жанр Драма есть в наборе")
	}

	// Удаляем жанр
	delete(genres, "Комедия")

	fmt.Println("Список жанров:")
	for genre := range genres {
		fmt.Println("-", genre)
	}

	item := CatalogItem{ID: 1, Title: "Интерстеллар", Price: 12.99, Available: true}
	fmt.Println(item)

	c1 := Coordinates{Lat: 55.75, Lng: 37.62}
	c2 := Coordinates{Lat: 55.75, Lng: 37.62}
	fmt.Println(c1 == c2) // true - все поля равны

	orig := Subscriber{Name: "Оригинал", Plan: "Basic"}
	copied := orig // создается копия
	copied.Name = "Копия"
	fmt.Println(orig.Name, copied.Name) // "Оригинал" "Копия"

	session := NewStreamSesseion("Владимир Путин", "Интепстеллар")
	fmt.Printf("%+v\n", *session)

}
