package main

type User struct {
	Id       string
	Name     string
	Presence Presence
	Channel  []Channel
}
