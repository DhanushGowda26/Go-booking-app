package main
import (
	"fmt"
	"time"
)
func sendTicket(userTicket uint,firstName,lastName,email string){
	var ticket =fmt.Sprintf("%v tickets for %v %v",userTicket,firstName,lastName)
	fmt.Println("#################")
	fmt.Println("Sending ticket...")
	time.Sleep(5*time.Second)
	fmt.Printf("sent ticket: %v to email adress %v\n",ticket,email)
	fmt.Println("#################")
	wg.Done()
}