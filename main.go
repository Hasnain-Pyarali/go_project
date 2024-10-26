package main

import (
	"fmt"
	"time"
	"sync"
)

var confName = "Go Conference"
const totalTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
type UserData struct {
	firstName string
	lastName string
	email string
	nbTickets uint
}

var wg = sync.WaitGroup {}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", totalTickets, remainingTickets)
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

	fmt.Println("Enter youre first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter youre last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter youre email adress:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		nbTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)
}

func sendTicket(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n%v\nto %v\n", ticket, email)
	fmt.Println("##########")
	wg.Done()
}

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(firstName, lastName, email, userTickets)
			fmt.Printf("These are all of our bookings %v\n", getFirstNames())
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Youre name or lastname is not valid")
			}
			if !isValidEmail {
				fmt.Println("Youre email is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets that you entered is not valid")
			}
		}
	}
	wg.Wait()
}
