package main

import (
	"fmt"
	"strings"
)

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
	
for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask for user input
		println("Enter your firstname!")
		fmt.Scan(&firstName)

		println("Enter your lastname!")
		fmt.Scan(&lastName)

		println("Enter your email address!")
		fmt.Scan(&email)

		println("Enter number of tickets!")
		fmt.Scan(&userTickets)

		isValidName := len(firstName)>=2 && len(lastName)>=2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber{

			//lets print the contents of the bookings array
			fmt.Printf("The whole slice : %v \n", bookings)
			fmt.Printf("The first value of the  slice : %v \n", bookings[0])
			fmt.Printf("The type of the slice : %T \n", bookings)
			fmt.Printf("The length of the slice : %v \n", len(bookings))

			fmt.Printf("Thank you  %v %v for booking %v tickets.\n You will receive a confirmation email at %v \n",firstName,lastName,userTickets,email)
			fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, element := range bookings{
				var names = strings.Fields(element) //Fileds is just like the split function in python. you pass in the element  accessed with the for loop. 
				//the fields method is from the strings library hence we need to import the string linrary
				firstNames = append(firstNames, names [0])
			}

			fmt.Printf("The first names of th bookings are : %v\n",firstNames)

			if remainingTickets == 0{

				//end program
				fmt.Println("Our Conference is booked out. Come back next year.")
				break
	
			}
				
		} else{
				fmt.Println("Your input data is invalid. Try again")
			}
		
		

		
		
		
	}
}
	