package main
import "fmt"

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

	fmt.Println("Enter youre first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter youre last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter youre email adress:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confName)
	fmt.Printf("These are all of our bookings %v\n", bookings)
}
