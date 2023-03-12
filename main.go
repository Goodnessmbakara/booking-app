package main

import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "our online booking application!")
	fmt.Println("we have a total of", conferenceTickets,"tickets and", remainingTickets, "are still available")
	fmt.Println(" pick up your tickets to attend.")
	
}