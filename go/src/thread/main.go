package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Printf("sdfghjkl ")
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		fmt.Printf("qwertyuiop ")
		time.Sleep(200 * time.Millisecond)
	}
}
