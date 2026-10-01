package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrUserNotFound = errors.New("пользователь не найден")

type ValidationError struct {
	Field string
	Value any
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("валидация поля: '%s': %s", e.Field, e.Value)
}

type User struct {
	ID string
}

type SubscritptionService struct {
	users map[string]User
}

func NewSubscriptionService() *SubscritptionService {
	return &SubscritptionService{
		users: map[string]User{
			"1": {ID: "1"},
			"2": {ID: "2"},
		},
	}
}

func (s *SubscritptionService) Activate(userID, dateStr string) error {
	if _, ok := s.users[userID]; !ok {
		return ErrUserNotFound
	}

	_, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return &ValidationError{
			Field: "date",
			Value: dateStr,
		}
	}

	fmt.Printf("Subscription activated for users %s until %s \n", userID, dateStr)
	return nil
}

func main() {
	svc := NewSubscriptionService()

	// Успешный случай
	err := svc.Activate("1", "2026-12-31")
	fmt.Println("1:", err) // <nil>

	// Пользователь не найден
	err = svc.Activate("999", "2026-12-31")
	fmt.Println("2:", err) // пользователь не найден

	// Неверная дата
	err = svc.Activate("1", "31-12-2026")
	fmt.Println("3:", err) // валидация поля: 'дата': 2026-12-31

	// Проверка через errors.Is
	err = svc.Activate("999", "2026-12-31")
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("User not found")
	}

	// Проверка через errors.As
	err = svc.Activate("1", "bad-date")
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Println("Field:", valErr.Field, "Value:", valErr.Value)
	}
}
