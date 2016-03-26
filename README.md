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
