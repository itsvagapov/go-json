package main

import (
	"encoding/json"
	"fmt"
	// "strings"
)

type School struct {
	school string
	location string
	groups []Group
}

type Group struct {
	title string
	location string
	remote bool
}

type WorkDay struct {
    Title       string   `json:"title"`
    City        string   `json:"city"`
    Address     string   `json:"address"`
    IsOnline    bool     `json:"is_online"`
    Seats       int      `json:"seats"`
    Mentors     []string `json:"mentors"`
}

type Student struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Age       int    `json:"age"`
    City      string `json:"city"`
    WantsGo   bool   `json:"wants_go"`
}

func main() {
	studentJSON := `{
  "first_name": "Иса",
  "last_name": "Исаев",
  "age": 17,
  "city": "Грозный",
  "wants_go": true
}`

	var s Student

	// fmt.Println(studentJSON)

	if err := json.Unmarshal([]byte(studentJSON), &s); err != nil {
		fmt.Println("Ошибка JSON:", err)
	}

	if s.WantsGo && s.Age >= 16 {
		fmt.Printf("%s %s готов к Буткемпу Intocode в Грозном!\n", s.FirstName, s.LastName)
	} else if s.Age < 16 {
		fmt.Println("слишком маленький возраст!")
	}
	

	// fmt.Println(s)

	// w := WorkDay{
		// Title: "Первый день в Intocode",
		// City: "Грозный",
		// Address: "ул. Дадин Айбики",
		// IsOnline: false,
		// Seats: 20,
		// Mentors: []string{"Кудузов Ахмад", "Назиров Расул"},
	// }

	// jsonBytes, _ := json.MarshalIndent(w, "", "\t")

	// fmt.Println(string(jsonBytes))
}