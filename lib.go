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
	ISBN string 	`json:"isbn,omitempty" gorm:"not null"`
}

func main() {
	db := connect()
	defer db.Close()
	fmt.Println("Welcome to Library System")
	fmt.Println("What do you want to do?\n1. Add a book\n2. Search a book or Delete\n3. Exit")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
		
	case 1:
		var rootCmd = &cobra.Command{Use: "app"}
		var createCmd = &cobra.Command{
			Use: "create",
			Short: "Create a new book",
			Run: func(cmd *cobra.Command, args []string){
				Title, _ := cmd.Flags().GetString("title")
				Author, _ := cmd.Flags().GetString("author")
				ISBN, _ := cmd.Flags().GetString("isbn")
				db := connect()
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
		createCmd.Flags().String("isbn", "", "ISBN of the book")
		createCmd.MarkFlagRequired("isbn")
		rootCmd.AddCommand(createCmd)
		rootCmd.Execute()

	case 2:
		fmt.Println("Search a book or Delete a book")
		var rootCmd = &cobra.Command{Use: "app"}
		var searchCmd = &cobra.Command{
			Use: "search",
			Run: func(cmd *cobra.Command, args []string){
				isbn, _ := cmd.Flags().GetString("isbn")
				var book Book
				db := connect()
				if err := db.Where("isbn = ?", isbn).First(&book).Error; err != nil {
					log.Fatal(err)
				}
				fmt.Println(book)
			},
		}
		var deleteCmd = &cobra.Command{
			Use: "delete",
			Run: func(cmd *cobra.Command, args []string){
				isbn, _ := cmd.Flags().GetString("isbn")
				var book Book
				db := connect()
				if err := db.Where("isbn = ?", isbn).First(&book).Error; err != nil {
					log.Fatal(err)
				}
				tx := connect().Begin()
				if err := tx.Delete(&book).Error; err != nil {
					tx.Rollback()
					return
				}
				tx.Commit()
				fmt.Println("Book deleted successfully")
			},
		}
		searchCmd.Flags().String("isbn", "", "ISBN of the book")
		searchCmd.MarkFlagRequired("isbn")
		deleteCmd.Flags().String("isbn", "", "ISBN of the book")
		deleteCmd.MarkFlagRequired("isbn")
		rootCmd.AddCommand(searchCmd, deleteCmd)
		rootCmd.Execute()

	case 3:
		fmt.Println("Thank you for using our system")

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
func connect() *gorm.DB{
	db, err := gorm.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&Book{}).Error; err != nil {
		log.Fatal(err)
	}
	return db
}