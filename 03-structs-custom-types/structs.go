package main

import (
	"fmt"
	"structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate) // creation / constructor func

	if err != nil {
		fmt.Println(err)
		return
	}

	adminData := user.NewAdmin("admin@mail.com", "123456", appUser.FirstName, appUser.LastName, appUser.Birthdate)

	appUser.OutputUserDetails()
	adminData.OutputAdminDetails()
	appUser.ClearUser()
	adminData.ClearUser()
	appUser.OutputUserDetails()
	adminData.OutputAdminDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
