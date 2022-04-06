package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // type inference
	const conferenceTickets = 50
	var remainingTickets uint = 50 // declare type
	bookings := []string{}

	// // print var types
	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.\n")

	for remainingTickets > 0 && len(bookings) <= 50 {
		// user input
		var firstName, lastName, email string
		var userTickets uint
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter the amount of tickets to purchase: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets // basic logic
			bookings = append(bookings, fmt.Sprintf("%v %v", firstName, lastName))

			// // arrays & slices
			// fmt.Printf("The whole slice: %v.\nThe first value: %v.\nSlice type: %T.\nSlice length: %v.\n", bookings, bookings[0], bookings, len(bookings))

			fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, strings.Fields(booking)[0])
			}
			fmt.Printf("These are the first name of our bookings: %v.\n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets!", remainingTickets, userTickets)
			continue
		}
	}
}
