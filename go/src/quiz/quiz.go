package quiz

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
	"math/rand"
	"reflect"
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

func RandItem(list []func() bool) func() bool {
	rand.Seed(t.Now().UnixNano())
	randNum := rand.Intn(len(list))
	randElement := list[randNum]
	return randElement
}

func Remove(slice []func() bool, item func() bool) []func() bool {
	var index int

	i2 := reflect.ValueOf(item)
	for p, v := range slice {
		v2 := reflect.ValueOf(v)
		if v2 == i2 {
			index = p
		}
	}
	var result []func() bool
	result = append(result[:index], result[index+1:]...)
	return result
}

func Questions() {
	questions := []func() bool{Q1, Q2, Q3}
	for questions != nil {
		question := RandItem(questions)
		if question() {
			questions = Remove(questions, question)
		}
	}
}
