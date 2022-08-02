package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// list of maps

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInputs()

	isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketsNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := printFirstNames()
		fmt.Printf("The first names of bookings are %v:\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out, Come back next year")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered id too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketsNumber {
			fmt.Println("number of tickets you entered is invalid")
		}

		fmt.Printf("Your input is invalid, please try again\n")
	}

	// city := "London"
	// switch city {
	// case "New York":
	// 	// Execute code to book tickets for New York
	// case "Singapore", "Hong Kong":
	// 	// Execute code to book tickets for Singapore
	// case "London", "Berlin":
	// 	// Execute code to book tickets for London
	// case "Mexico City":
	// 	// Execute code to book tickets for Mexico
	// default:
	// 	fmt.Println("No Valid City Selected")

	// }

	wg.Wait()

}

func greetUsers() {

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for their name
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets you want")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}
