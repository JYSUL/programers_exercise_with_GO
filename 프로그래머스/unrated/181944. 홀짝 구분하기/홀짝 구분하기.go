package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    if a % 2 == 1 {
        fmt.Printf("%v is odd\n",a)
    } else {
        fmt.Printf("%v is even\n",a)
    }
}