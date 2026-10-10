package main

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

func main() {

}
