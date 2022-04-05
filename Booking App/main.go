package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %s booking application. \nGet your tickets here to attend.\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}
