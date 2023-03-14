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
	//fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T", conferenceName, conferenceTickets, remainingTickets)

	greetUsers(conferenceName,conferenceTickets,remainingTickets)

	// we want t0 create an array to store bookings

	var bookings =[]string{}
	
	for {
			firstName,lastName,email,userTickets := getUserInput()

			isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName,lastName,userTickets,remainingTickets,email)
			if isValidEmail && isValidName && isValidTicketNumber{

				//lets print the contents of the bookings array
				fmt.Printf("The whole slice : %v \n", bookings)
				fmt.Printf("The first value of the  slice : %v \n", bookings[0])
				fmt.Printf("The type of the slice : %T \n", bookings)
				fmt.Printf("The length of the slice : %v \n", len(bookings))
				
				bookTicket(remainingTickets , userTickets, bookings, firstName , lastName, conferenceName )

				firstNames := getFirstNames(bookings)
				fmt.Printf("the first names of the bookings are : %v",firstNames)
				if remainingTickets == 0{

					//end program
					fmt.Println("Our Conference is booked out. Come back next year.")
					break
		
				}
					
			} else{
					if !isValidName{
						fmt.Println("the first name or last name you entered is too short")
					}
					if !isValidEmail{
						fmt.Println("The email you entered is incorrect")
					}
					if !isValidTicketNumber{
						fmt.Println("the number of tickets you entered is invalid")
					}
					fmt.Println("Your input data is invalid. Try again")
				}	
		}
	}

func greetUsers(conferenceName string,conferenceTickets uint,remainingTickets uint){
	fmt.Printf("Welcome to %v our online booking application!\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Printf("pick up your tickets to attend.\n")		
	}

func getFirstNames(bookings[] string )[] string{
	firstNames := []string{}
				for _, element := range bookings{
					var names = strings.Fields(element) //Fileds is just like the split function in python. you pass in the element  accessed with the for loop. 
					//the fields method is from the strings library hence we need to import the string linrary
					firstNames = append(firstNames, names [0])
				}

				return firstNames
}

func validateUserInput(firstName string, lastName string, userTickets uint, remainingTickets uint, email string)(bool,bool,bool){
	isValidName := len(firstName)>=2 && len(lastName)>=2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicketNumber
}
	

func getUserInput()(string,string,string,uint){
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

	return firstName,lastName,email,userTickets
}


func bookTicket(remainingTickets uint, userTickets uint, bookings [] string, firstName string, lastName string, conferenceName string){
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + "" + lastName)
	
	fmt.Printf("Thank you %v %v for booking %v Tickets. You rock!\n",firstName,lastName,userTickets)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)	

}