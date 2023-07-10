package main
import("fmt"
"strings")
func validateInput(firstName, lastName, email, city *string, userTicket *uint) {
	for {
		fmt.Printf("Please enter your first name: ")
		fmt.Scanln(firstName)
		isValidFName := len(*firstName) >= 2
		if !isValidFName {
			fmt.Printf("First name entered is too short\n")
			continue
		}
		fmt.Printf("Please enter your last name: ")
		fmt.Scanln(lastName)
		isValidLName := len(*lastName) >= 2
		if !isValidLName {
			fmt.Printf("Last name entered is too short\n")
			continue
		}
		fmt.Printf("Please enter your email: ")
		fmt.Scanln(email)
		isValidEmail := strings.Contains(*email, "@") && strings.Contains(*email, ".")
		if !isValidEmail {
			fmt.Printf("Please enter a valid email\n")
			continue
		}
		fmt.Printf("Please enter city (Bengaluru or Mumbai): ")
		fmt.Scanln(city)
		isValidCity := *city == "Bengaluru" || *city == "Mumbai"
		if !isValidCity {
			fmt.Printf("You entered an invalid city\n")
			continue
		}
		fmt.Printf("Hello %v %v, how many tickets would you like to buy in %v? ", *firstName, *lastName, *city)
		fmt.Scanln(userTicket)
		isValidTicket := *userTicket > 0 && *userTicket <= remainingTickets
		if !isValidTicket {
			fmt.Printf("Number of tickets you entered is invalid. Out of %v, %v tickets are remaining\n", confrenceTicket,remainingTickets)
			continue
		}
		break
	}
}