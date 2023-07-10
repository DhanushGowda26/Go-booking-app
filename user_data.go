package main
// import "strings"
func getBookings() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// names := strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}