package main

import "fmt"

func main() {
	var conferenceName = "Go Con"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and there are are %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

}
