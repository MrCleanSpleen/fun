package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	t "time"
)

func main() {
	const snake = `
	          /\ /\         /\ /\         /\ /\              _______   
                 / / \ \       / / \ \       / / \ \             \   _  \  
                / /   \ \     / /   \ \     / /   \ \            /  /_\  \ 
                / /    \ \    / /    \ \    / /    \ \           |   _    \
               / /     \ \   / /     \ \   / /     \ \           \  \_/    \ 
 ____________ / /       \ \ / /       \ \ / /       \ \___________\_____   /
/_____/_____/ \/         \/ \/         \/ \/         \/_____/_____/     \_/ 
	`
	for {
		for i := 0; i < 4; i++ {
			for i := 0; i < 12; i++ {
				t.Sleep(300 * t.Millisecond)
				fmt.Println(c.B03 + "Badgers" + c.X)
			}
			fmt.Println(c.CL)
			for i := 0; i < 2; i++ {
				fmt.Println(c.R + "Mushroom!")
				t.Sleep(750 * t.Millisecond)
			}
			fmt.Println(c.CL)
		}
		for i := 0; i < 1; i++ {
			fmt.Print(c.X + snake)
			fmt.Println(c.X + "Ahhh! A Snake!")
			t.Sleep(5500 * t.Millisecond)
		}
	}
}
