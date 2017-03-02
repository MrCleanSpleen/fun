package main
import (
	"os"
	"fmt"
)
cur1 := os.Args[1]
amount := os.Args[2]
cur2 := os.Args[3]
url := fmt.Sprintf("http://api.fixer.io/latest")
type Rates struct {
	Cur1	float64	`json:

