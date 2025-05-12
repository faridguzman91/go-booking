package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "go conf"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("welcome to %v booking app \n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v still available \n", conferenceTickets, remainingTickets)
	fmt.Println("get your ticket here")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("enter first name:")
		fmt.Scan(&firstName)

		fmt.Println("enter last name:")

		fmt.Scan(&lastName)

		fmt.Println("enter email:")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets:")
		fmt.Scan(&userTickets)

    if userTickets > remainingTickets {
      fmt.Printf("We only have %v tickets remaining, so cant book %v tickets\n", remainingTickets, userTickets)
      //skip all other executions, continue
      continue
    }
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("thanks for booking %v %v you booked %v tickets. \n you should be getting an email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		//firstname slice
		firstNames := []string{}

		// for loop (index number and booking in bookings, look in range on data structure of bookings)
		// end when iterating iver all bookings elements, _ to define unused variable (index will not be used)

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

    noTicketsRemaining := remainingTickets == 0  

    if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked ut, come back next year")
			break
		}

	}
}
