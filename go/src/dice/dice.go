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
 ------
|      |
|  壹  |
|     1|
 ------
	`)
		side2 = (`
 ------
|      |
|  貳  |
|     2|
 ------
	`)

		side3 = (`
 ------
|      |
|  叁  |
|     3|
 ------
	`)

		side4 = (`
 ------
|      |
|  肆  |
|     4|
 ------
	`)

		side5 = (`
 ------
|      | 
|  伍  |
|     5|
 ------
	`)

		side6 = (`
 ------
|      | 
|  陸  |
|     6|
 ------
	`)
	)
	sides := []string{side1, side2, side3, side4, side5, side6}
	for {
		side := r.Strings(sides)
		fmt.Print(c.Clear + c.Rc() + side + c.X)
		i.Ask("")
	}
}
