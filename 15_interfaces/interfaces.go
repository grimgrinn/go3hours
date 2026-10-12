package main

import "fmt"

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
