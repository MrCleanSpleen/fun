package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

func main() {
	for {
		fmt.Print(c.Rc() + "Nyan " + c.X)
	}
}
