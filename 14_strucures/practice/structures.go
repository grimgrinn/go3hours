package main

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	UserName string
	Email    string
	Age      int
	Country  string
	Premium  bool
}

type Device struct {
	Name   string
	Type   string
	Active bool
}

func (d Device) Info() string {
	return fmt.Sprintf("%s, %s, %v", d.Name, d.Type, d.Active)
}

func (d *Device) Activate() {
	d.Active = true
}

func (d *Device) Deactivate() {
	d.Active = false
}

type Subscription struct {
	Plan      string
	Price     float64
	AutoRenew bool
}

func NewSubcription(plan string, price float64) *Subscription {
	return &Subscription{
		Plan:      plan,
		Price:     price,
		AutoRenew: true,
	}
}

func (s *Subscription) CancelAutorenew() {
	s.AutoRenew = false
}

func (s Subscription) PrintDetails() string {
	return fmt.Sprintf("Plan: %s, Price: %.1f, Autorenew: %v", s.Plan, s.Price, s.AutoRenew)
}

type Episode struct {
	Title    string `json:"title"`
	Season   int    `json:"season"`
	Duration int    `json:"duration"`
}

type WatchHistory struct {
	Profile         Profile
	Episode         Episode
	WatchedPercent  int
	LastPositionMin int
}

type FavoriteList struct {
	Titles []string
	uniq   map[string]struct{}
}

func (f *FavoriteList) Add(title string) {
	f.Titles = append(f.Titles, title)
}

func (f FavoriteList) PrintAll() {
	fmt.Println()
	for _, title := range f.Titles {
		fmt.Println(title)
	}
}

func (f *FavoriteList) PrintUnique() {
	if f.uniq == nil {
		f.uniq = make(map[string]struct{})
	}
	fmt.Println()
	for _, title := range f.Titles {
		f.uniq[title] = struct{}{}
	}

	for utitle := range f.uniq {
		fmt.Println(utitle)
	}
}

func main() {
	user := Profile{UserName: "Владимир Путин", Email: "putin@kremlin.ru", Age: 51, Country: "Russia", Premium: true}
	fmt.Println(user)

	var user2 Profile
	fmt.Println(user2)

	user3 := Profile{UserName: "Арнольд Шварцнегер", Email: "schwarzenegger@schwarznegger.com"}
	fmt.Println(user3)

	device := Device{Name: "Magnitofone", Type: "Stereo", Active: false}

	fmt.Println(device.Info())
	device.Activate()
	fmt.Println(device.Info())
	device.Deactivate()
	fmt.Println(device.Info())

	sub := NewSubcription("premium", 100)
	fmt.Println(sub.PrintDetails())
	sub.CancelAutorenew()
	fmt.Println(sub.PrintDetails())

	episode := Episode{"First", 3, 190}
	historyEntry := WatchHistory{
		Profile:         user3,
		Episode:         episode,
		WatchedPercent:  90,
		LastPositionMin: 170,
	}
	fmt.Println(historyEntry.Profile.UserName, historyEntry.Profile.Email, historyEntry.Episode.Title, historyEntry.WatchedPercent, historyEntry.LastPositionMin)
	b, err := json.Marshal(historyEntry)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	fl := FavoriteList{Titles: []string{"Terminatror", "Rocky", "Saw", "Inception"}}
	fl.PrintAll()
	fl.Add("Terminator")
	fl.Add("Terminator")
	fl.PrintAll()
	fl.Add("Terminator 2")
	fl.PrintAll()
	fl.PrintUnique()

	temp := struct {
		Owner      string
		ItemsCount int
	}{
		Owner:      "Putin",
		ItemsCount: 12039,
	}
	fmt.Println(temp)

}
