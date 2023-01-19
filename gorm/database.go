package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookId int
	BookName string
	AuthorName string
	BookPrice int
	BookQuantity int
}

func Database(){
	db,err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}else {
		fmt.Println("Database connected successfully")
	}
	db.AutoMigrate(&Book{})
}