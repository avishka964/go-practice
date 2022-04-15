package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and remaning %v tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email:")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want:")
		fmt.Scan(&userTickets)

		if userTickets >= int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("There are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out")
				break
			}
		} else {
			fmt.Println("Wrong input")
		}

	}

}
