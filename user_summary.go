package main

import ("fmt")

type UserData struct{
		firstName string
		lastName string	
		email string
		city string
		noOfTickets uint
	
	}
func summary(userTicket *uint, firstName, lastName, city,email *string) {
	// to print confirmation message
	if *userTicket <= remainingTickets {
		remainingTickets = remainingTickets - *userTicket
		

		var userData = UserData{
				firstName: *firstName,
				lastName: *lastName,
				email: *email,
				city: *city,
				noOfTickets: *userTicket,
		}
		bookings = append(bookings, userData)
		fmt.Printf("Booking Summary: %v\n",userData)
		fmt.Printf("Thank You %v %v, You have booked %v ticket(s) for %v in %v. You will receive a confirmation email at %v\n", *firstName, *lastName, *userTicket, confrenceName, *city, *email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, confrenceName)
		firstNames := getBookings()
		fmt.Printf("Members booked are %v\n", firstNames)
	}
}