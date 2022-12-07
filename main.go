package main

import "fmt"

func main() {
	const DaysTotal = 90
	var remainingDays uint = 90
	challenge := "#90DaysOfDevops"

	fmt.Printf("welcome to the %v challenge.\n This challenge consists of %v days\n", challenge, DaysTotal)

	var TwitterName string
	var DaysCompleted uint

	fmt.Println("Enter your Twitter Handle: ")
	fmt.Scanln(&TwitterName)

	fmt.Println("How Many days have you completed ")
	fmt.Scanln(&DaysCompleted)

	remainingDays = remainingDays - DaysCompleted

	fmt.Printf("Thank you %v for taking and completing %v days. ", TwitterName, DaysCompleted)
	fmt.Printf("You have %v days remaining for the %v challenge\n", remainingDays, challenge)
	fmt.Println("Good Luck !")
}
