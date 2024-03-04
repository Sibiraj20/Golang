package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	// Ask user for name
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)

	// Ask user for number of tickets
	fmt.Print("How many tickets would you like to book? ")
	fmt.Scanln(&userTickets)

	if userTickets <= remainingTickets {
		fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
		remainingTickets -= userTickets
		fmt.Printf("%v tickets are now remaining.\n", remainingTickets)
	} else {
		fmt.Println("Sorry, we don't have enough tickets available for your request.")
	}
}

