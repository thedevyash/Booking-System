package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcom to our", conferenceName, "booking application")
	fmt.Println("We have", conferenceTickets, "tickets available", "and these many are remaining", remainingTickets)
	fmt.Println("Get your ticket now")

	var userName string
	//this is a comment
	userName = "Tom"
	fmt.Printf("hello %v", userName)

}
