package helper	
import "fmt"

func GetUserInput()(string,string,string,uint){
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