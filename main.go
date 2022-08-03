package main

import "fmt"

func main() {
	// var conferenceName = "Go Con" --can't define constants like this, or an explicit data type
	conferenceName := "Go Con"
	const conferenceTickets = 50
	var remainingTickets uint = 50 // uint allows whole numbers only, no negative values

	fmt.Printf("conferenceName is type %T, conferenceTickets is type %T, remainingTickets is type %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and there are are %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName) // the & means a pointer to a place in memory - doesn't exist in Java
	// so what happens is Scan looks at what there in memory and scans it
	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets // guessing -= works here

	fmt.Printf("Thank you, %v %v, for booking %v tickets.  You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
