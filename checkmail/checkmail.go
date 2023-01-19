package checkmail

import (
	"fmt"
	"github.com/badoux/checkmail"
)

func Check(){
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