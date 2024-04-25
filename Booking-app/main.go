package main

import (
	"Booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference" //conferenceName:="Go Conference" -- only used for variables , cant assign variable type
const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings = [50]string{}
// var bookings = make([]map[string]string, 0)
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetusers()

	// fmt.Println("Welcome to", conferenceName, "booking application!")

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidemail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidTicketNumber && isValidemail && isValidName {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of th bookings are: %v\n", firstNames)

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
	wg.Wait()
}

func greetusers() {
	fmt.Printf("Welcome to %v booking application\n.", conferenceName)
	fmt.Printf("We have total of %v, tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName

	//create map for the user
	// var userData = make(map[string]string) //we can't mix different data types
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	// fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets.You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	println("########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	println("########################")
	wg.Done()
}
