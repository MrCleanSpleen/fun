package quiz

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
	"strings"
)

func Yes() {
	fmt.Println(c.M + "Huzzah! You are correct!")
}
func No() {
	fmt.Println(c.R + "I am sorry, you are incorrect.")
}
func Q1() {
	a1, err := i.Prompt(c.M + "What color are the unicorns? ")
	a1
	if strings.Contains(a1, "pink") {
		Yes()
	} else {
		No()
	}
}
func Q2() {
	a2, err := i.Prompt(c.M + "Where are the dancing? ")
	a2
	if strings.Contains(a2, "rainbows") {
		Yes()
	} else {
		No()
	}
}

func Q3() {
	a3, err := i.Prompt(c.M + "Please use 1 word to describe the texture of their maaaaaaaaaaaagical fur. ")
	a3
	if strings.Contains(a3, "smiles") {
		Yes()
	} else {
		No()
	}
}
func Questions() {
	Q1()
	Q2()
	Q3()

}
