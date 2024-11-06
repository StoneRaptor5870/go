package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age:", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(*agePointer)
}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
