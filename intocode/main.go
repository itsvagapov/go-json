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

func main() {
	w := WorkDay{
		Title: "Первый день в Intocode",
		City: "Грозный",
		Address: "ул. Дадин Айбики",
		IsOnline: false,
		Seats: 20,
		Mentors: []string{"Кудузов Ахмад", "Назиров Расул"},
	}

	jsonBytes, _ := json.MarshalIndent(w, "", "\t")

	fmt.Println(string(jsonBytes))
}