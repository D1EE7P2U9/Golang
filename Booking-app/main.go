package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference" //conferenceName:="Go Conference" -- only used for variables , cant assign variable type
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	// var bookings = [50]string{}
	var bookings []string

	// fmt.Println("Welcome to", conferenceName, "booking application!")
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total of %v, tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their Name
		fmt.Println("Enter your first Name:")
		fmt.Scan(&firstName) //& is used for pointers ,
		// fmt.Println(conferenceName)
		// fmt.Println(&conferenceName)  -- prints memory location of the var conferenceName
		fmt.Println("Enter your last Name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		//input validation
		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		isValidemail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidTicketNumber && isValidemail && isValidName {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets.You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first Names of bookings are: %v\n", firstNames)
			if remainingTickets == 0 {
				//end
				fmt.Println("Our conference is booked out. Come next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("firstName or lastName you entered is too short")
			}
			if !isValidemail {
				fmt.Println("email you entered doesn't contain @ sign")
			}
			if isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}
}
