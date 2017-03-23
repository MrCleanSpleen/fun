package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test := []string{
		"1234567890",
		"qwertyuiop",
		"asdfghjkl",
		"asdfjasdgfasgf",
		"zxctrfvbnjhbnmj",
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(test)
	fmt.Println(test[n])
}
