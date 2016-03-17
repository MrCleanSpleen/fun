package dice
import (
	"fmt"
	c"github.com/skilstak/go/colors"
	r"github.com/skilstak/go/choice"
	i"github.com/skilstak/go/input"
)
func Dice() {
	sides := []
	const sides = ('''
 -----
|     |
|  0  |
|     |
 -----
	''')

	const sides = ('''
 -----
|    0|
|     |
|0    |
 -----
	''')

	const sides = ('''
 -----
|    0|
|  0  |
|0    |
 -----
	''')	

	const sides = ('''
 -----
|0   0|
|     |
|0   0|
 -----
	''')

	const sides = ('''
 -----
|0   0| 
|  0  |
|0   0|
 -----
	''')

	const sides = ('''
 -----
|0 0 0| 
|     |
|0 0 0|
 -----
	''')

	for {
		side := r.Strings(sides)
		fmt.Print(c.Clear + c.Rc() + side + c.X)
		i.Ask()
	}
}
