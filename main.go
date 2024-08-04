package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go conference"
	//or    conferenceName:="Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	fmt.Println("Welcom to our", conferenceName, "booking application")
	fmt.Println("We have", conferenceTickets, "tickets available", "and these many are remaining", remainingTickets)
	fmt.Println("Get your ticket now")
	//dynamic list or slice
	var bookings []string
	//or bookings:=[]string{}

	for {
		var firstName string
		var lastName string
		// bookings[0] = firstName + " " + lastName

		var email string
		var userTickets uint
		//this is a comment
		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)
		var isValidName = len(firstName) > 0 && len(lastName) > 0

		bookings = append(bookings, firstName+" "+lastName)
		fmt.Println("Please enter your email")
		fmt.Scan(&email)
		var isValideEmail = strings.Contains(email, "@") && strings.Contains(email, ".")
		fmt.Println("Please enter the number of tickets you want to book")
		fmt.Scan(&userTickets)
		validTicketsnumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValideEmail && validTicketsnumber {
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("Thank you %v for booking %v tickets for our %v", firstName, userTickets, conferenceName)
			fmt.Printf("We have %v tickets remaining\n", remainingTickets)
			firstNames := []string{}
			for _, booking := range bookings {
				//string.Fields gives a list of words/strings seperated by word as seperate elements of an array
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The person who have booked the ticket are %v\n", firstNames)
		}

		if remainingTickets == 0 {
			fmt.Println("We are sold out")
			break

		}
	}
}
