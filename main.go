package main
import (
	"fmt"
	"strings"
)

func main() {
	var confName = "Go Conference"
	const totalTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		fmt.Println("Enter youre first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter youre last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter youre email adress:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(firstName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName + " " + lastName)
	
			firstNames := []string{}
	
			for _, booking := range bookings {
				firstNames = append(firstNames, strings.Fields(booking)[0])
			}
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)
			fmt.Printf("These are all of our bookings %v\n", firstNames)
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

}
