package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

// package level variables → accessible by all functions everywhere
var conferenceName string = "Go Conference" // type inference
var remainingTickets uint = 50              // declare type
var bookings = []string{}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application. \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets left.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Fields(booking)[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter the amount of tickets to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets // basic logic
	bookings = append(bookings, fmt.Sprintf("%v %v", firstName, lastName))

	// // arrays & slices
	// fmt.Printf("The whole slice: %v.\nThe first value: %v.\nSlice type: %T.\nSlice length: %v.\n", bookings, bookings[0], bookings, len(bookings))

	fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func main() {

	// // print var types
	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	greetUsers()

	for remainingTickets > 0 && len(bookings) <= 50 {

		// user input
		firstName, lastName, email, userTickets := getUserInput()

		// input validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are the first name of our bookings: %v.\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("First name or last name entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email does not contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid amount of tickets.")
			}

		}

	}

	// // switch statement
	// city := "London"
	// switch city {
	// case "New York":
	// 	// execute code for New York
	// case "Singapore":
	// 	// execute code for Singapore
	// case "London", "Berlin", "Paris":
	// 	// execute for Europe
	// default:
	// 	fmt.Println("No valid city selected.")
	// }

}
