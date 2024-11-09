package anonymous

import (
	"fmt"
)

func Anon() {
	// anonymous function
	numbers := []int{1, 2, 3}
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// closures
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
