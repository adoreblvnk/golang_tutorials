package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // type inference
	const conferenceTickets = 50
	var remainingTickets uint = 50 // declare type

	// print var types
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}
