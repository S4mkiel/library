package main

import (
	"fmt"
	"os"
)




func main() {
	fmt.Println("Welcome to Library System")
	fmt.Println("What do you want to do?\n1. Add a book\n2. Remove a book\n3. Search a book\n4. Exit")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Add a book")

		erro, file:= os.OpenFile("books.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		fmt.Println("Enter the book name:")
		var bookName string
		fmt.Scanln(&bookName)

		fmt.Println("Enter the author name:")
		var authorName string
		fmt.Scanln(&authorName)

		fmt.Println("Enter the book price:")
		var bookPrice int
		fmt.Scanln(&bookPrice)

		fmt.Println("Enter the book quantity:")
		var bookQuantity int
		fmt.Scanln(&bookQuantity)

		fmt.Println("Book added successfully")

	case 2:
		fmt.Println("Remove a book")
	case 3:
		fmt.Println("Search a book")
	case 4:
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid choice")
	}

}