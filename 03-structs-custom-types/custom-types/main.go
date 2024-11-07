package main

import "fmt"

type str string // custom type str of string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Nischay"
	name.log()
}
