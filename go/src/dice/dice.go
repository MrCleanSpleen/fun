package dice

import (
	"fmt"
	r "github.com/skilstak/go/choice"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
)

func Dice() {
	const (
		side1 = (`
 -----
|     |
|  0  |
|     |
 -----
	`)
		side2 = (`
 -----
|    0|
|     |
|0    |
 -----
	`)

		side3 = (`
 -----
|    0|
|  0  |
|0    |
 -----
	`)

		side4 = (`
 -----
|0   0|
|     |
|0   0|
 -----
	`)

		side5 = (`
 -----
|0   0| 
|  0  |
|0   0|
 -----
	`)

		side6 = (`
 -----
|0 0 0| 
|     |
|0 0 0|
 -----
	`)
	)
	sides := []string{side1, side2, side3, side4, side5, side6}
	for {
		side := r.Strings(sides)
		fmt.Print(c.Clear + c.Rc() + side + c.X)
		i.Ask("")
	}
}
