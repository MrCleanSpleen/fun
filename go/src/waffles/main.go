package main

import (
	"fmt"
	i "github.com/skilstak/go/input"
)

func main() {
	waffles := i.Ask("Do you like waffles? >>> ")
	if waffles == "yes" {
		fmt.Println(":)")
		pancakes := i.Ask("Do you like pancakes? >>> ")
		if pancakes == "yes" {
			fmt.Println(":)")
			toast := i.Ask("Do you like french toast? >>> ")
			if toast == "yes" {
				fmt.Println("Do do dah do, can't wait to get a mouthful!")

			} else {
				fmt.Println("")
			}
		} else {
			fmt.Println("")
		}
	} else {
		fmt.Println("")
	}
}
