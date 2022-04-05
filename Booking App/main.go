package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // type inference
	const conferenceTickets = 50
	var remainingTickets uint = 50 // declare type
	var bookings [50]string

	// // print var types
	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// user input
	var username, email string
	var userTickets uint
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter the amount of tickets to purchase: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets // basic logic
	bookings[0] = username

	// arrays & slices
	fmt.Printf("The whole array: %v.\nThe first value: %v.\nArray type: %T.\n Array length: %v.\n", bookings, bookings[0], bookings, len(bookings))

	fmt.Printf("Thank you %v for purchasing %v tickets. You will receive a confirmation email at %v.\n", username, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
