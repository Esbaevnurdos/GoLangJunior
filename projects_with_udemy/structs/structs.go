package main

import (
	"fmt"

	"example.com/structs/user"
)



func main() {
    userFirstName := getUserInput("First name: ")
	userLastName := getUserInput("Last name: ")
	userBirthDate := getUserInput("Birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
       fmt.Println(err)
	   return
	}

	admin := user.NewAdmin("test@gmail.com", "test123")

	admin.User.OutputUserInput()
	admin.User.ClearUserName()
	admin.User.OutputUserInput()

	appUser.OutputUserInput()
	appUser.ClearUserName()
	appUser.OutputUserInput()
}



func getUserInput(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}