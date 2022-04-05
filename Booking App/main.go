package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // type inference
	const conferenceTickets = 50
	var remainingTickets uint = 50 // declare type

	// // print var types
	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// user input
	var username, email string
	var userTickets int
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter the amount of tickets to purchase: ")
	fmt.Scan(&userTickets)
	fmt.Printf("Thank you %v for purchasing %v tickets. You will receive a confirmation email at %v.\n", username, userTickets, email)
}
