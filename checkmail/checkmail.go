package checkmail

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"github.com/badoux/checkmail"
)

func Login(){
	fmt.Println("Welcome to System")
	fmt.Println("What do you want to do?\n1. Login\n2. Register\n3. Exit")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Write your email address")
		var email string
		fmt.Scanln(&email)
		err := checkmail.ValidateFormat(email)
		if err != nil {
			fmt.Println("Invalid email address")
		}else {
			fmt.Println("Valid email address")
		}

		fmt.Println("Write your password")
		var password string
		fmt.Scanln(&password)
		hash := sha256.New()
		hash.Write([]byte(password))
	case 2:
		Register()
	}
}

func Register(){
	fmt.Println("Please enter register details")
	regEmail()
	regPassword()
	file, err := os.OpenFile("users.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil{
		log.Fatal(err)
	} else {
		file.WriteString("email")
		file.WriteString("password")
	}
	if file != nil {
		fmt.Println("User registered successfully")
	}else {
		fmt.Println("User registration failed")
	}
	defer file.Close()
	for {
		fmt.Println("Do you want to login? (y/n)")
		var choice string
		fmt.Scanln(&choice)
		if choice == "y"{
			Login()
			break
		} else if choice == "n"{
			fmt.Println("Thank you for using our system")
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}
}

func regEmail(){
	fmt.Println("Write your email address")
	var email string
	fmt.Scanln(&email)
	err := checkmail.ValidateFormat(email)
	if err != nil {
		fmt.Println("Invalid email address")
	}else {
		fmt.Println("Valid email address")
	}
}

func regPassword(){
	var password string
	fmt.Println("Write your password")
	fmt.Scanln(&password)
	if len(password) < 8 {
		fmt.Println("Password must be at least 8 characters")
	}else if len(password) > 20 {
		fmt.Println("Password must be less than 20 characters")
	}else {
		fmt.Println("Valid password")
	}
	hash := sha256.New()
	hash.Write([]byte (password))
	fmt.Printf("%x", hash.Sum(nil))
}