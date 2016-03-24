package quiz

import (
	"fmt"
	r "github.com/skilstak/go/choice"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
	"strings"
)

func Yes() {
	fmt.Println("Huzzah! You are correct!")
}
func No() {
	fmt.Println("I am sorry, you are incorrect.")
}
func Q1() {
	a1 := i.Prompt("What color are the unicorns? ")
	a1
	if strings.Contains(a1, "pink") {
		Yes()
	} else {
		No()
	}
}
func Q2() {
	a2 := i.Prompt("Where are the dancing? ")
	a2
	if strings.Contains(a2, "rainbows") {
		Yes()
	} else {
		No()
	}
}

func Q3() {
	a3 := i.Prompt("Please use 1 word to descrive the texture of their maaaaaaaaaaaagical fur. ")
	a3
	if strings.Contains(a3, "smiles") {
		Yes()
	} else {
		No()
	}
}
func Quiz() {
}
