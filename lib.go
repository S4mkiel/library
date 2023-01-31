package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/cobra"
)

type Book struct {
	gorm.Model
	Title  string `json:"title,omitempty" gorm:"not null"`
	Author string `json:"author,omitempty" gorm:"not null"`
	ISBN uint 		`json:"isbn,omitempty" gorm:"not null"`
}

func main() {
	fmt.Println("Welcome to Library System")
	fmt.Println("What do you want to do?\n1. Add a book\n2. Search a book or Delete\n3. Exit")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Add a book")
		db, err := gorm.Open("sqlite3", "library.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		if err := db.AutoMigrate(&Book{}).Error; err != nil {
			log.Fatal(err)
		}
		var rootCmd = &cobra.Command{Use: "app"}
		var createCmd = &cobra.Command{
			Use: "create",
			Run: func(cmd *cobra.Command, args []string){
				Title, _ := cmd.Flags().GetString("title")
				Author, _ := cmd.Flags().GetString("author")
				ISBN, _ := cmd.Flags().GetUint("isbn")
				tx := db.Begin()
				if err := tx.Create(&Book{Title: Title, Author: Author, ISBN: ISBN}).Error; err != nil {
					tx.Rollback()
					return
				}
				tx.Commit()
				fmt.Println("Book added successfully")
			},
		}
		createCmd.Flags().String("title", "", "Title of the book")
		createCmd.MarkFlagRequired("title")
		createCmd.Flags().String("author", "", "Author of the book")
		createCmd.MarkFlagRequired("author")
		createCmd.Flags().Uint("isbn", 0, "ISBN of the book")
		createCmd.MarkFlagRequired("isbn")
		rootCmd.AddCommand(createCmd)
		rootCmd.Execute()
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