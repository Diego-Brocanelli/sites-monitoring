package options

import "fmt"

const lineSeparator = "------------------------------------------------------------------"

// Sow options
func ShowOptions() {
	fmt.Println(lineSeparator)
	fmt.Println("Welcome to the site monitoring system, choose the option you want.")
	fmt.Println(lineSeparator)
	fmt.Println("Option: 0 - Exit System")
	fmt.Println("Option: 1 - Start Monitoring")
	fmt.Println("Option: 2 - Show logs")
}

// Responsible for quantity for monitoring
func chooeseQuantity() {
	fmt.Println(lineSeparator)
	fmt.Println("Choose the amount of analysis you want.")
	fmt.Println(lineSeparator)
}

// Get command chosen by the user
func GetCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

// Get quantity chosen by the user
func GetQuantity() int {

	chooeseQuantity()

	var quantity int
	fmt.Scan(&quantity)

	return quantity
}
