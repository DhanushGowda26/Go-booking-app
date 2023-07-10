package greet
import ("fmt")
func GreetUsers(confName string,confTicket int,remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n",confName)
	fmt.Printf("Total %v tickets available and %v tickets remaining\n",confTicket,remainingTickets)
	fmt.Printf("Get your tickets to attend\n")
}