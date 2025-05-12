package main

import "fmt"

func main() {
	conferenceName := "go conf"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("welcome to %v booking app \n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v still available \n", conferenceTickets, remainingTickets)
	fmt.Println("get your ticket here")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	var bookings []string

	// ask user for name (entr in console)
	fmt.Println("enter first name:")
	fmt.Scan(&firstName)

	fmt.Println("enter last name:")

	fmt.Scan(&lastName)

	fmt.Println("enter email:")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

  // fmt.Printf("The whole array: %v\n", bookings)
  // fmt.Printf("The first value: %v\n", bookings[0])
  // fmt.Printf("Slice type: %T\n", bookings)
  // fmt.Printf("Slice length: %v\n", len(bookings))


	fmt.Printf("thanks for booking %v %v you booked %v tickets. \n you should be getting an email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

  fmt.Printf("These are all our bookings: %v\n", bookings)
}
