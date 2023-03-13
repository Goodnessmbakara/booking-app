package main

import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	//lets print out the types of variables we have
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v our online booking application!\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Printf("pick up your tickets to attend.\n")


	var firstName string
	var lastName string
	var email string
	var userTickets int

	//ask for user input
	println("Enter your firstname!")
	fmt.Scan(&firstName)

	println("Enter your lastname!")
	fmt.Scan(&lastName)

	println("Enter your email address!")
	fmt.Scan(&email)

	println("Enter number of tickets!")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you  %v %v for booking %v tickets. You will receive a confirmation email at %v \n",firstName,lastName,userTickets,email)
	remainingTickets  = uint(remainingTickets) - uint(userTickets)

	fmt.Printf("%v tickets remaining for the %v", remainingTickets, conferenceName)



	
	
}