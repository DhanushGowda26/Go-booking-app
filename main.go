package main
import( "booking-app/greetings"
	"sync")
var wg =sync.WaitGroup{
	}
const confrenceTicket int = 50
var confrenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
func main() {
	for{
		// Initializing variables
		var firstName string
		var lastName string
		var email string
		var city string
		var userTicket uint
		// greeting user
		greet.GreetUsers(confrenceName , confrenceTicket, remainingTickets)
		// aking user to provide input and validating the input
		validateInput(&firstName, &lastName, &email, &city, &userTicket)
		// Update the remaining tickets and process booking
		summary(&userTicket, &firstName, &lastName,&city ,&email)
		// Check if tickets are sold out
		if soldOut() {
			break // Exit the loop if sold out
		}
		// Add a wait group to send the ticket
		wg.Add(1)
		// confirm booking
		go sendTicket(userTicket, firstName, lastName, email)	
	wg.Wait()
}
}