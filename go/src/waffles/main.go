package main

import (
	"fmt"
	// c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
)

func main() {
	waffles := i.Ask("Do you like Waffles? ")
	if waffles == "yes" {
		pancakes := i.Ask("Do you like pancakes? ")
		if pancakes == "yes" {
			french_toast := i.Ask("Do you like french toast? ")
			if french_toast == "yes" {
				fmt.Print("Do do dah do, can't wait to get a mouthful!")
			} else {
				fmt.Print("Now you shall DIIEEEEEEEEEEEEE!")
			}

		} else {
			fmt.Print("Now you shall DIIEEEEEEEEEEEEE!")
		}

	} else {
		fmt.Print("Now you shall DIIEEEEEEEEEEEEE!")
	}
}
