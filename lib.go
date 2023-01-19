package main

import (
	"fmt"
	"library/checkmail"
)

func main() {
	checkmail.Login()
	fmt.Println("Welcome to Library System")
	fmt.Println("What do you want to do?\n1. Add a book\n2. Search a book or Delete\n3. Exit")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		var bookName, authorName  string
		var bookQuantity, bookPrice int
		fmt.Println("Add a book")

		fmt.Println("Enter the book name:")
		fmt.Scanln(&bookName)

		fmt.Println("Enter the author name:")
		fmt.Scanln(&authorName)

		fmt.Println("Enter the book price:")
		fmt.Scanln(&bookPrice)

		fmt.Println("Enter the book quantity:")
		fmt.Scanln(&bookQuantity)

		fmt.Println("Book added successfully")
		
	case 2:
		fmt.Println("Search a book or Delete a book")
					
	case 3:
		fmt.Println("Exit")
		for {
			fmt.Println("Are you sure you want to exit? (y/n)")
			var choice string
			fmt.Scanln(&choice)
			if choice == "y" {
				fmt.Println("Thank you for using our system")
				break
			} else if choice == "n" {
				fmt.Println("Welcome back")
				main()
				break
			} else {
				fmt.Println("Invalid choice")
			}
		}
	default:
		fmt.Println("Invalid choice")
		for {
			fmt.Println("Do you want to try again? (y/n)")
			var choice string
			fmt.Scanln(&choice)
			if choice == "y" {
				main()
				break
			} else if choice == "n" {
				fmt.Println("Thank you for using our system")
				break
			} else {
				fmt.Println("Invalid choice")
			}
		}
	}
}