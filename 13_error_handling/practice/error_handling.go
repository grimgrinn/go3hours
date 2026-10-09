package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Movie struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func loadCatalog(files []string) (movies []Movie, err error) {
	start := time.Now()

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	defer func() {
		fmt.Printf("loadCatalog took: %v\n", time.Since(start))
	}()

	for _, path := range files {
		fileMovies, readErr := readMoviesFile(path)
		if readErr != nil {
			return nil, fmt.Errorf("read file %s: %w", path, readErr)
		}
		movies = append(movies, fileMovies...)
	}

	return movies, nil
}

func readMoviesFile(path string) ([]Movie, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer f.Close()

	var movies []Movie

	if err := json.NewDecoder(f).Decode(&movies); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	return movies, nil
}

func main() {
	movies, err := loadCatalog([]string{"movies1.json", "movies2.json"})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Loaded %d movies\n", len(movies))
	for _, m := range movies {
		fmt.Printf("- %s (%d)\n", m.Name, m.Year)
	}
}
