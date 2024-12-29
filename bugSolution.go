```go
package main

import "fmt"

func main() {
    var x int = 10
    var y *int = &x
    *y = 20
    fmt.Println(x) // Output: 20

    var z int = 30
    // Create a new pointer to z
    var w *int = &z
    *w = 40
    fmt.Println(x, z) // Output: 20, 40
}
```