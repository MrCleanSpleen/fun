package main

import (
	"fmt"
)

func main() {
	fmt.Println("What is your name?")
	var input string
	fmt.Scanln(&input)
	fmt.Print("Nice to meet you " + input + "!")
}
