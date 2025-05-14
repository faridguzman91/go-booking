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

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {

    firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("thanks for booking %v %v you booked %v tickets. \n you should be getting an email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			printFirstnames(bookings)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked ut, come back next year")
				break
			}

		} else if userTickets == remainingTickets {
			//do something
		} else {
			//fmt.Printf("We only have %v tickets remaining, so cant book %v tickets\n", remainingTickets, userTickets)
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email ddress is not valid, does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets is invalid")
			}
			fmt.Printf("your data is invalid, try again")
		}
	}
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Println("Welcome to the conf")
	fmt.Printf("welcome to %v booking app \n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v still available \n", conferenceTickets, remainingTickets)
	fmt.Println("get your ticket here")
}

func printFirstnames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber

}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}
