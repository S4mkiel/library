package main

import (
	"fmt"
	"log"
	"strings"

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
	fmt.Println("Welcome to Library System")
	fmt.Println("What do you want to do?\n1. Add a book\n2. Search a book or Delete\n3. Exit")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Write the details of the book\nPs: Title, Author and ISBN are required with command: create --title <title> --author <author> --isbn <isbn>")
		var create string
		fmt.Scanln(&create)

		createArgs := strings.Split(create, " ")
		var rootCmd = &cobra.Command{Use: "app"}
		var createCmd = &cobra.Command{
			Use:   "create",
			Short: "Create a new book",
			Run: func(cmd *cobra.Command, args []string) {
				title, _ := cmd.Flags().GetString("title")
				author, _ := cmd.Flags().GetString("author")
				isbn, _ := cmd.Flags().GetString("isbn")
				db := connect()
				defer db.Close()
				if err := db.Create(&Book{Title: title, Author: author, ISBN: isbn}).Error; err != nil {
					log.Fatalf("Failed to create book: %v", err)
				}
				fmt.Println("Book added successfully")
			},
		}
		createCmd.Flags().StringP("title", "t", "", "title of the book")
		createCmd.Flags().StringP("author", "a", "", "author of the book")
		createCmd.Flags().StringP("isbn", "i", "", "isbn of the book")
		rootCmd.AddCommand(createCmd)
		rootCmd.SetArgs(createArgs)
		err := rootCmd.Execute()
		if err != nil {
			fmt.Println(err)
		}
	case 2:
		fmt.Println("Search a book or Delete a book\nPS: Search a book with command: search --title <title> and Delete a book with command: delete --isbn <isbn>")
		var search string
		fmt.Print("Enter search term: ")
		fmt.Scanln(&search)
		searchArgs := strings.Split(search, " ")

		var rootCmd = &cobra.Command{Use: "app"}
		var searchCmd = &cobra.Command{
			Use:   "search",
			Short: "Search for books in the database",
			Run: func(cmd *cobra.Command, args []string) {
				title, _ := cmd.Flags().GetString("title")
				var books []Book
				db := connect()
				if err := db.Where("title LIKE ?", "%"+title+"%").Find(&books).Error; err != nil {
					log.Fatal(err)
				}
				for _, book := range books {
					fmt.Println(book)
				}
			},
		}
		var deleteCmd = &cobra.Command{
			Use: "delete",
			Short: "Delete a book from the database",
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
		var Title string
		searchCmd.Flags().StringVarP(&Title, "title", "t", "", "title of the book")
		deleteCmd.Flags().String("isbn", "", "ISBN of the book")
		deleteCmd.MarkFlagRequired("isbn")
		rootCmd.AddCommand(searchCmd, deleteCmd)
		rootCmd.SetArgs(searchArgs)
		if err := rootCmd.Execute(); err != nil {
			log.Fatal(err)
		}
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