package main

import (
	"fmt"
	"functions/anonymous"
	"functions/functionsarevalue"
	"functions/recursion"
	"functions/variadicfunc"
)

func main() {
	fmt.Println("-------Functions as value--------")
	functionsarevalue.FuncsAsValues()
	fmt.Println("--------Anonymous Functions & Closures--------")
	anonymous.Anon()
	fmt.Println("--------Recursion--------")
	recursion.Recursion()
	fmt.Println("--------Variadic Function--------")
	variadicfunc.Variadic()
}
