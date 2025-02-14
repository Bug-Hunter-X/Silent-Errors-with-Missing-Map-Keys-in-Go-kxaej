```go
package main

import "fmt"

func main() {
    var m = make(map[int]int)
    m[1] = 10

    value, ok := m[2] // Check if key exists using the comma ok idiom
    if ok {
        fmt.Println("Value found for key 2:", value)
    } else {
        fmt.Println("Key 2 not found")
    }

    value, ok = m[100]
    if ok {
        fmt.Println("Value found for key 100:", value)
    } else {
        fmt.Println("Key 100 not found")
    }
}
```