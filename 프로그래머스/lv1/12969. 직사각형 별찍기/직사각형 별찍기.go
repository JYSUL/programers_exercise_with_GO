package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    var printer string = ""
    for i := 0; i < a; i++ {
        printer += "*"
    }
    for i := 0; i < b; i++ {
        fmt.Println(printer)
    }
}