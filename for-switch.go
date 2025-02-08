package main

import "fmt"

func main() {
	for {
		var choice int
		fmt.Print("Enter a number (1-3) or 0 to exit: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("You chose one.")
		case 2:
			fmt.Println("You chose two.")
		case 3:
			fmt.Println("You chose three.")
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option.")

		}
	}
}
