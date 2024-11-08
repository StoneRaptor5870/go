package main

import (
	"fmt"
	"map/maps"
)

type floatMap map[string]float64 // type alias

func main() {
	maps.Map()
	fmt.Println("------------------------------")

	userName := make([]string, 2, 5) // type, size, capacity
	userName[0] = "Nischay"
	userName = append(userName, "Max", "John")
	fmt.Println(userName)
	fmt.Println("------------------------------")

	skills := make(floatMap, 5)
	skills["JS"] = 4.7
	skills["TS"] = 4.8
	skills["React"] = 4.5
	fmt.Println(skills)
	fmt.Println("------------------------------")

	// loops for arrays, slice & map
	for index, value := range userName {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}
	fmt.Println("------------------------------")

	for key, value := range skills {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
