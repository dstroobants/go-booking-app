package helper

import (
	"fmt"
	"strconv"
	"strings"
)

// Capitalize first letter of function to export it to other packages
func ValidateUserInput(remainingTickets uint, firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func GreetUsers(remainingTickets uint, confName string, confTickets int) {
	fmt.Printf("Welcome to the %v booking application\n", confName)
	fmt.Printf("There is a total of %v tickets, %v tickets are still available\n", confTickets,
		remainingTickets)
	fmt.Println("Get your tickets here to attend the event")
}

func GetFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	// iterate over all elements using blank identifier (Ignore index)
	for _, booking := range bookings {
		// append the first name to the slice
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func GetUserInput() (string, string, string, uint) {
	// If we do not assign a value to a variable, we need to specify the type
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask user for their input
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Please enter the number of tickets you would like to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func BookTicket(remainingTickets uint, userTickets uint, bookings []map[string]string, firstName string, lastName string, email string, conferenceName string) []map[string]string {
	remainingTickets = remainingTickets - userTickets

	// create map for user data
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Println("There are", remainingTickets, "tickets remaining")
	return bookings
}
