package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	createdAt time.Time
}

type Admin struct { // embedding / inheritance
	email    string
	password string
	User
}

func (u User) OutputUserDetails() { // method of the struct by adding (u user)
	fmt.Println(u.FirstName, u.LastName, u.Birthdate)
}

func (a Admin) OutputAdminDetails() { // method of the struct by adding (a Admin)
	fmt.Println(a.email, a.password, a.FirstName, a.LastName, a.Birthdate)
}

func (u *User) ClearUser() { // mutation method - pass as pointers
	u.FirstName = ""
	u.LastName = ""
}

func NewAdmin(email, password, firstName, lastName, birthdate string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: firstName,
			LastName:  lastName,
			Birthdate: birthdate,
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) { // creation / constructor func
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("fields cant be left blank")
	}

	return &User{ // returning pointer
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
