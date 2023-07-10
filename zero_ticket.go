package main
import "fmt"
var soldOutCh = make(chan struct{}) // Channel to signal sold-out status
func soldOut() bool {
	if remainingTickets == 0 {
		fmt.Printf("%v is sold out!\n", confrenceName)
		close(soldOutCh) // Signal sold-out status
		return true
	}
	return false
}