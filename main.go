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



	// we want t0 create an array to store bookings

	var bookings =[]string{}
	

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


	remainingTickets  = uint(remainingTickets) - uint(userTickets)
	bookings = append (bookings, firstName + " " + lastName)

	//lets print the contents of the bookings array
	fmt.Printf("The whole slice : %v \n", bookings)
	fmt.Printf("The first value of the  slice : %v \n", bookings[0])
	fmt.Printf("The type of the slice : %T \n", bookings)
	fmt.Printf("The length of the slice : %v \n", len(bookings))

	fmt.Printf("Thank you  %v %v for booking %v tickets.\n You will receive a confirmation email at %v \n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)

	fmt.Printf("these are all our bookings : %v \n",bookings)



	
	
}