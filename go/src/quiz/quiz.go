package quiz

import (
	"fmt"
	r "github.com/skilstak/go-random"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
	"strings"
	t "time"
)

func Yes() {
	fmt.Println(c.CL + c.M + "Huzzah! You are correct!")
	t.Sleep(2 * t.Second)
}
func No() {
	fmt.Println(c.CL + c.R + "I am sorry, you are incorrect.")
	t.Sleep(2 * t.Second)
}
func Q1() bool {
	a1 := i.Ask(c.CL + c.M + "What color are the unicorns? \n")
	if strings.Contains(a1, "pink") {
		Yes()
		return true
	}
	No()
	return false

}
func Q2() bool {
	a2 := i.Ask(c.CL + c.M + "Where are the dancing? \n")
	if strings.Contains(a2, "rainbows") {
		Yes()
		return true
	}
	No()
	return false
}

func Q3() bool {
	a3 := i.Ask(c.CL + c.M + "Please use 1 word to describe the texture of their maaaaaaaaaaaagical fur. \n")
	if strings.Contains(a3, "smiles") {
		Yes()
		return true
	}
	No()
	return false
}
func Questions() {
	questions := []func(){Q1, Q2, Q3}
	quiz := (questions)

}
