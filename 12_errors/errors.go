package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidUser = errors.New("неверный пользователь")
var ErrMovieNotFound = errors.New("Фильм не найден")
var ErrAccessDenied = errors.New("доступ запрещен")

type ValidationError struct {
	Field   string
	Value   any
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("валидация поля '%s': %s", e.Field, e.Message)
}

func loadMovie(title string) (string, error) {
	catalog := map[string]string{
		"Интерстеллар": "Фантастика, 2014, 8.6",
		"Начало":       "Фантастика, 2010, 8.8Э",
	}

	data, ok := catalog[title]
	if !ok {
		return "", fmt.Errorf("фильм '%s' не найден", title)
	}
	return data, nil
}

func processSubscription(userName, plan string) (string, error) {
	if userName == "" {
		return "", ErrInvalidUser
	}

	validPlans := map[string]bool{
		"Basic":    true,
		"Standart": true,
		"Premium":  true,
	}

	if !validPlans[plan] {
		return "", fmt.Errorf("недопустимый план: '%s'", plan)
	}

	return fmt.Sprintf("%s подключил план '%s'", userName, plan), nil
}

func findMovieById(id int) (string, error) {
	switch id {
	case 42:
		return "", ErrMovieNotFound
	case 100:
		return "", ErrAccessDenied
	default:
		return fmt.Sprintf("Фильм #%d", id), nil
	}
}

func watchMovie(movieID int) error {
	title, err := findMovieById(movieID)
	if err != nil {
		return fmt.Errorf("просмотр фильма %d: %w", movieID, err)
	}
	fmt.Printf("> %s\n", title)
	return nil
}

func validateUserAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "не может быть отрицательным",
		}
	}

	if age > 100 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "нереалистичный возраст",
		}
	}

	return nil
}

func startMovieStream(userName string, movieID int) error {
	if userName == "" {
		return fmt.Errorf("авторизация: %w", ErrInvalidUser)
	}

	title, err := findMovieById(movieID)
	if err != nil {
		return fmt.Errorf("загрузка: %w", err)
	}

	fmt.Printf("> %s запускает '%s'\n", userName, title)
	return nil
}

func main() {
	fmt.Println("=== Патттерн result, err ===")

	movie, err := loadMovie("Интерстеллар")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Загружен:", movie)
	}

	movie, err = loadMovie("НесуществущийФильм")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("=== if c инициализацией ===")

	if movie, err := loadMovie("Начало"); err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Загружен:", movie)
	}

	fmt.Println("=== Early return ===")

	result, err := processSubscription("Алексей", "Premium")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("=== Оборачивание ошибок (%w) ===")

	err = watchMovie(42)
	if err != nil {
		fmt.Println("Ошибка:", err)

		if errors.Is(err, ErrMovieNotFound) {
			fmt.Println("> Фмльм не найден")
		}
	}

	fmt.Println("\n=== Сигнальные (sentinel) ошибки ===")

	for _, id := range []int{1, 42, 100} {
		title, err := findMovieById(id)

		switch {
		case errors.Is(err, ErrMovieNotFound):
			fmt.Printf("ID %d: фильм не найден\n", id)
		case errors.Is(err, ErrAccessDenied):
			fmt.Printf("ID %d: доступ запрещен\n", id)
		case err != nil:
			fmt.Printf("ID %d: %v\n", id, err)
		default:
			fmt.Printf("ID %d: %s\n", id, title)
		}
	}

	fmt.Println("\n=== Игнорирование ошибки ===")

	value, _ := strconv.Atoi("123")
	fmt.Println("Значение:", value)

	fmt.Println("=== errors.As ===")

	err = validateUserAge(-9)
	if err != nil {
		var validErr *ValidationError

		if errors.As(err, &validErr) {
			fmt.Printf("Поле '%s' : %s [значение: %v]\n",
				validErr.Field, validErr.Message, validErr.Value)
		}
	}

	fmt.Println("\n=== Цепочка обработок ===")

	if err := startMovieStream("Алексей", 1); err != nil {
		fmt.Println("Стрим:", err)
	}

	if err := startMovieStream("Алексей", 42); err != nil {
		fmt.Println("Стрим:", err)
	}

	if err := startMovieStream("", 1); err != nil {
		fmt.Println("Стрим:", err)
	}
}
