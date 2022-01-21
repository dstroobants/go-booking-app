package main

import (
	"fmt"
	"go-booking-app/helper"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers(conferenceName, conferenceTickets)

	for { // Infinite loop

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(remainingTickets, firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber { // If user tries to book more tickets than are available

			bookTicket(userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names in the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out")
				break
			}
		} else {
			switch {
			case !isValidName:
				fmt.Println("Please enter a valid name")
			case !isValidEmail:
				fmt.Println("Please enter a valid email")
			case !isValidTicketNumber:
				fmt.Println("Please enter a valid ticket number")
			}
		}
	}
	fmt.Println("Exiting booking application...")

}

func greetUsers(confName string, confTickets int) {
	fmt.Printf("Welcome to the %v booking application\n", confName)
	fmt.Printf("There is a total of %v tickets, %v tickets are still available\n", confTickets,
		remainingTickets)
	fmt.Println("Get your tickets here to attend the event")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	// iterate over all elements using blank identifier (Ignore index)
	for _, booking := range bookings {
		// split the string and return a slice of split values
		// "Nicole Smith" => ["Nicole", "Smith"]
		var names = strings.Fields(booking)
		var firstName = names[0]
		// append the first name to the slice
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

func bookTicket(userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	fmt.Printf(
		"Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Println("There are", remainingTickets, "tickets remaining")
}

/*

	//var conferenceName = "Go conference"
	const conferenceTickets = 50 // value cannot be changed
	//var remainingTickets = 50

	// alternative way of declaring variables, cannot be used for constants
	conferenceName := "Go conference"
	remainingTickets := 50

	//var bookings = [50] string {"Alex", "Steve"} // array of strings with two default values

	// Slices can be used to create dynamic arrays (lists)
	var bookings [] string
	// alternative syntax for slice creation
	bookings := [] string {}

	//fmt.Print("Hello World") // Prints inline
	//fmt.Println("Hello World") // Prints on new line

	//fmt.Println("Welcome to the", conferenceName, "booking application")
	// Same as above but with formatting
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)

	//fmt.Println("There is a total of", conferenceTickets, "tickets,", remainingTickets,
		//"tickets are still available")
	// Same as above but with formatting
	fmt.Printf("There is a total of %v tickets, %v tickets are still available\n", conferenceTickets,
		remainingTickets)

	fmt.Println("Get your tickets here to attend the event")

	// If we do not assign a value to a variable, we need to specify the type
	var firstName string
	var lastName string
	var email string
	var userTickets int

	//firstName = "Denis"
	//userTickets = 2

	//Pointers
	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets) // Variable containing the memory address of remainingTickets

	//fmt.Println("Hello", firstName, "you have", userTickets, "tickets")

	// Print the types of variables
	fmt.Printf("conferenceName type is %T and conferenceTickets type is %T\n", conferenceName,
		conferenceTickets)

	// Ask user for their input
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Please enter the number of tickets you would like to purchase: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	//bookings[2]= firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)

	//fmt.Printf("The whole array is %v\n", bookings)
	//fmt.Printf("Array type is %T\n", bookings)
	//fmt.Printf("Array length is %v\n", len(bookings))

	fmt.Printf("The whole slice is %v\n", bookings)
	fmt.Printf("slice type is %T\n", bookings)
	fmt.Printf("slice length is %v\n", len(bookings))

	fmt.Println("There are", remainingTickets, "tickets remaining")

*/

//----------------------------------------------------------------------------------------------
// Clean up 1
//----------------------------------------------------------------------------------------------

/*
			conferenceName := "Go conference"
			const conferenceTickets = 50
			remainingTickets := 50
			bookings := []string{}

			fmt.Printf("Welcome to the %v booking application\n", conferenceName)
			fmt.Printf("There is a total of %v tickets, %v tickets are still available\n", conferenceTickets,
				remainingTickets)
			fmt.Println("Get your tickets here to attend the event")

			for { // Infinite loop
				// If we do not assign a value to a variable, we need to specify the type
				var firstName string
				var lastName string
				var email string
				var userTickets int

				// Ask user for their input
				fmt.Println("Please enter your first name: ")
				fmt.Scan(&firstName)
				fmt.Println("Please enter your last name: ")
				fmt.Scan(&lastName)
				fmt.Println("Please enter your email: ")
				fmt.Scan(&email)
				fmt.Println("Please enter the number of tickets you would like to purchase: ")
				fmt.Scan(&userTickets)

				if userTickets > remainingTickets { // If user tries to book more tickets than are available
					fmt.Println("We only have %v tickets remaining\n", remainingTickets)
					continue // Go back to the beginning of the loop
				}

				remainingTickets = remainingTickets - userTickets
				bookings = append(bookings, firstName+" "+lastName)

				fmt.Printf(
					"Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n",
					firstName, lastName, userTickets, email)

				firstNames := []string{}
				// iterate over all elements using blank identifier (Ignore index)
				for _, booking := range bookings {
					// split the string and return a slice of split values
					// "Nicole Smith" => ["Nicole", "Smith"]
					var names = strings.Fields(booking)
					var firstName = names[0]
					// append the first name to the slice
					firstNames = append(firstNames, firstName)

				}
				fmt.Println("The first names in the bookings are: %v\n", firstNames)
				fmt.Println("There are", remainingTickets, "tickets remaining")

				if remainingTickets == 0 {
					fmt.Println("Sorry, we are sold out")
					break
				}
			}

			fmt.Println("Exiting booking application...")

		//----------------------------------------------------------------------------------------------
		// Clean up 2
		//----------------------------------------------------------------------------------------------

		conferenceName := "Go conference"
		const conferenceTickets = 50
		remainingTickets := 50
		bookings := []string{}

		fmt.Printf("Welcome to the %v booking application\n", conferenceName)
		fmt.Printf("There is a total of %v tickets, %v tickets are still available\n", conferenceTickets,
			remainingTickets)
		fmt.Println("Get your tickets here to attend the event")

		for { // Infinite loop
			// If we do not assign a value to a variable, we need to specify the type
			var firstName string
			var lastName string
			var email string
			var userTickets int

			// Ask user for their input
			fmt.Println("Please enter your first name: ")
			fmt.Scan(&firstName)
			fmt.Println("Please enter your last name: ")
			fmt.Scan(&lastName)
			fmt.Println("Please enter your email: ")
			fmt.Scan(&email)
			fmt.Println("Please enter the number of tickets you would like to purchase: ")
			fmt.Scan(&userTickets)

			if userTickets <= remainingTickets { // If user tries to book more tickets than are available
				remainingTickets = remainingTickets - userTickets
				bookings = append(bookings, firstName+" "+lastName)

				fmt.Printf(
					"Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n",
					firstName, lastName, userTickets, email)

				firstNames := []string{}
				// iterate over all elements using blank identifier (Ignore index)
				for _, booking := range bookings {
					// split the string and return a slice of split values
					// "Nicole Smith" => ["Nicole", "Smith"]
					var names = strings.Fields(booking)
					var firstName = names[0]
					// append the first name to the slice
					firstNames = append(firstNames, firstName)

				}
				fmt.Println("The first names in the bookings are: %v\n", firstNames)
				fmt.Println("There are", remainingTickets, "tickets remaining")

				if remainingTickets == 0 {
					fmt.Println("Sorry, we are sold out")
					break
				}
			} else {
				fmt.Println("We only have %v tickets remaining\n", remainingTickets)
			}
		}
		fmt.Println("Exiting booking application...")
		// END




	//----------------------------------------------------------------------------------------------
	// User input validation
	//----------------------------------------------------------------------------------------------

	conferenceName := "Go conference"
	const conferenceTickets = 50
	remainingTickets := 50
	bookings := []string{}

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("There is a total of %v tickets, %v tickets are still available\n", conferenceTickets,
		remainingTickets)
	fmt.Println("Get your tickets here to attend the event")

	for { // Infinite loop
		// If we do not assign a value to a variable, we need to specify the type
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// Ask user for their input
		fmt.Println("Please enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Please enter the number of tickets you would like to purchase: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		isValidCity := city == "Singapore" || city == "London"

		if userTickets <= remainingTickets { // If user tries to book more tickets than are available
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf(
				"Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v\n",
				firstName, lastName, userTickets, email)

			firstNames := []string{}
			// iterate over all elements using blank identifier (Ignore index)
			for _, booking := range bookings {
				// split the string and return a slice of split values
				// "Nicole Smith" => ["Nicole", "Smith"]
				var names = strings.Fields(booking)
				var firstName = names[0]
				// append the first name to the slice
				firstNames = append(firstNames, firstName)

			}
			fmt.Println("The first names in the bookings are: %v\n", firstNames)
			fmt.Println("There are", remainingTickets, "tickets remaining")

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out")
				break
			}
		} else {
			fmt.Println("We only have %v tickets remaining\n", remainingTickets)
		}
	}
	fmt.Println("Exiting booking application...")

*/
//----------------------------------------------------------------------------------------------
// Functions
//----------------------------------------------------------------------------------------------
