package gorm 

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	
)

type Book struct {
	gorm.Model
	BookName string
	AuthorName string
	BookPrice int
	BookQuantity int
}

func main (){
	//...
}