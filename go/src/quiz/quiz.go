package quiz

import (
	"fmt"
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
func Q1() {
	a1 := i.Ask(c.CL + c.M + "What color are the unicorns? ")
	if strings.Contains(a1, "pink") {
		Yes()
	} else {
		No()
	}
}
func Q2() {
	a2 := i.Ask(c.CL + c.M + "Where are the dancing? ")
	if strings.Contains(a2, "rainbows") {
		Yes()
	} else {
		No()
	}
}

func Q3() {
	a3 := i.Ask(c.CL + c.M + "Please use 1 word to describe the texture of their maaaaaaaaaaaagical fur. ")
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
