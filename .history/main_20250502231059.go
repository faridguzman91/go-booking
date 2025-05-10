package main import "fmt" func main() { conferenceName := "go conf" const conferenceTickets = 50 var remainingTickets = 50 fmt.Printf("welcome to %v booking app \n", conferenceName) fmt.Printf(
    "we have a total of %v tickets and %v still available \n",
    conferenceTickets,
    remainingTickets
) fmt.Println("get your ticket here") var firstName string var lastName string var email string var userTickets int // ask user for name (entr in console) fmt.Println("enter first name:") fmt.Scan(& firstName) fmt.Println("enter last name:") fmt.Scan(& lastName) fmt.Println("enter email:") fmt.Scan(& email) fmt.Println("enter number of tickets:") fmt.Scan(& userTickets) fmt.Printf(
    "thanks for booking %v %v you booked %v tickets. \n you should be getting an email at %v",
    firstName,
    lastName,
    userTickets,
    email
) }