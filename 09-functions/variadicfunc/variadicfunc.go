package variadicfunc

import "fmt"

func Variadic() {
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 45, -5)
	anotherSum := sumup(numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
