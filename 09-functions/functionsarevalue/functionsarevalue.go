package functionsarevalue

import (
	"fmt"
)

// functions as values and function type

type transformFn func(int) int

func FuncsAsValues() {
	numbers := []int{1, 2, 3, 4}
	nums := []int{2, 3, 4, 5}
	doubled := transformNumbers(&numbers, getTransformerFunction(&numbers))
	tripled := transformNumbers(&nums, getTransformerFunction(&nums))
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}

	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}

	return tNumbers
}

// returning functions as value
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
