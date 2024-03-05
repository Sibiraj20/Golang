package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Define conference details
	var conferenceName = "Go Conference"
	const conferenceCapacity = 100
	var remainingTickets = conferenceCapacity

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total capacity of %v tickets\n", conferenceCapacity)
	fmt.Println("Get your tickets here to attend")

	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	for {
		var userName string
		var userTickets int

		// Ask user for name
		fmt.Print("Enter your name (or 'exit' to quit): ")
		fmt.Scanln(&userName)

		if userName == "exit" {
			fmt.Println("Exiting booking application.")
			break
		}

		// Ask user for number of tickets
		fmt.Print("How many tickets would you like to book? ")
		fmt.Scanln(&userTickets)

		if userTickets <= remainingTickets && userTickets > 0 {
			fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
			remainingTickets -= userTickets
			fmt.Printf("%v tickets are now remaining.\n", remainingTickets)

			// Simulate payment process (random success/failure)
			if rand.Intn(2) == 0 {
				fmt.Println("Payment successful! Tickets booked.")
			} else {
				fmt.Println("Payment failed. Please try again.")
				remainingTickets += userTickets // Return tickets if payment failed
			}
		} else if userTickets <= 0 {
			fmt.Println("Invalid number of tickets. Please try again.")
		} else {
			fmt.Println("Sorry, we don't have enough tickets available for your request.")
		}

		fmt.Println("---------------------------------------------")
	}

	fmt.Println("Thank you for using our booking application.")
}
