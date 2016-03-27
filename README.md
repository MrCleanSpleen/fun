# fun
The fundamentals of golang, python, and any other language I learn.

Nyan in Golang

1.0
```
package main
import (
  "fmt"
  c "github.com/skilstak/go/colors"
)
```
2.0
```
package main
import (
    "fmt"
    c"github.com/skilstak/go/colors"
)

func main() {
    fmt.Println("Nyan")
}
```
3.0
```

package main

import (
    "fmt"
    c "github.com/skilstak/go/colors"
)

func main() {
    fmt.Print(c.Rc() + "Nyan " + c.X)
}
```
4.0
```
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
```

Waffles in Golang

1.0
```
package main

import (
    "fmt"
    i "github.com/skilstak/go/input"
)

func main() {

    waffles := i.Ask("Do you like Waffles? ")
}
```
2.0
```
package main

import (
    "fmt"
    i "github.com/skilstak/go/input"
)

func main() {
    waffles := i.Ask("Do you like Waffles? ")
    if waffles == "yes" {
        fmt.Println("♪ Do do dah do, can't wait to get a mouthfull! ♪")
    }
}
```
3.0
```
package main

import (
    "fmt"
    i "github.com/skilstak/go/input"
)

func main() {
    waffles := i.Ask("Do you like Waffles? ")
    if waffles == "yes" {
        fmt.Println("♪ Do do dah do, can't wait to get a mouthfull! ♪")
    } else {
          fmt.Println("Humph. :(")
    }
}
```
4.0
```
package main

import (
    "fmt"
    i "github.com/skilstak/go/input"
)

func main() {
    waffles := i.Ask("Do you like Waffles? ")
    if waffles == "yes" {
      pancakes := i.Ask("Do you like pancakes? ")
      if pancakes == "yes" {
          fmt.Println("♪ Do do dah do, can't wait to get mouthful ♪")    
  }

      } else {
          
          fmt.Println("Humph. :(")
      } else {
          fmt.Println("Humph. :(")
    }
}
```
5.0
```
package main

import (
  "fmt"
  // c "github.com/skilstak/go/colors"
  i "github.com/skilstak/go/input"
)

func main() {
      waffles := i.Ask("Do you like Waffles? ")
      if waffles == "yes" {
        } pancakes := i.Ask("Do you like pancakes? ")
            if pancakes == "yes" {
              } french_toast := i.Ask("Do you like french toast? ")
                    if french_toast == "yes" {
                        fmt.Print("Do do dah do, can't wait to get a mouthful!")
                  } else {
                          fmt.Println("Humph. :(")
          } else {
                fmt.Println("Humph. :(")
    } else {
  
          fmt.Println("Humph. :(")
      }
}
```
