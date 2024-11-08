package main

import (
	"fmt"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	productNames[2] = "carpet"

	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[1:] // slice 1:3 first element is included and 3 is excluded
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// slices can modify the original values of the array
	// capacity is based on the items on the left the length of whixh the slice is based on

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	// Dynamic lists with slices

	costs := []float64{10.99, 8.99}
	fmt.Println(costs[0:1])
	costs[1] = 9.99

	updatedCosts := append(costs, 5.99)
	fmt.Println(costs)
	fmt.Println(updatedCosts)

	// unpacking lists or spreading
	discountPrices := []float64{2.5, 5.6, 9.5}
	costs = append(costs, discountPrices...)
	fmt.Println(costs)

	// practice
	fmt.Println("--------------------------------------------")

	hobbies := [4]string{"music", "games", "cooking", "travel"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	fmt.Println(cap(mainHobbies))

	mainHobbies = mainHobbies[1:4]
	fmt.Println(mainHobbies)

	courseGoals := []string{"learn go", "learn basics"}
	fmt.Println(courseGoals)
	courseGoals[1] = "learn how to build backends in go"
	courseGoals = append(courseGoals, "build cli tools")
	fmt.Println(courseGoals)

	products := []Product{{"1", "p1", 5.99}, {"2", "p2", 10.50}}
	fmt.Println(products)

	newProduct := Product{"3", "p3", 12.35}
	products = append(products, newProduct)
	fmt.Println(products)
}
