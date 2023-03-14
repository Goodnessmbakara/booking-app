package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]userData,0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	//lets print out the types of variables we have
	//fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T", conferenceName, conferenceTickets, remainingTickets)

	greetUsers()

	// we want t0 create an array to store bookings

	for {
		firstName, lastName, email, userTickets := helper.GetUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userTickets, email)
		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("the first names of the bookings are : %v", firstNames)
			if remainingTickets == 0 {

				//end program
				fmt.Println("Our Conference is booked out. Come back next year.")
				break

			}

		} else {
			if !isValidName {
				fmt.Println("the first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is incorrect")
			}
			if !isValidTicketNumber {
				fmt.Println("the number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v our online booking application!\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("pick up your tickets to attend.\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings{
		// var names = strings.Fields(element) //Fileds is just like the split function in python. you pass in the element  accessed with the for loop.
		//the fields method is from the strings library hence we need to import the string linrary

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, userTickets uint, email string) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicketNumber
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create map for user
	//map is a collection of key value pairs

	var userData userData = userData {
		firstName:firstName,
		lastName:        lastName,
		 numberOfTickets: userTickets,
		email:           email,
	}

	bookings = append(bookings, userData)
	fmt.Printf("The whole map content : %v \n", bookings)

	fmt.Printf("The length of the map : %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v Tickets. You rock!\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
