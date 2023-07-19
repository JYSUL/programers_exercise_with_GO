package main

import(
    "fmt"
    "strconv"
)

func main() {
    var s1 string
    var s2 string
    var p2 *string = &s2
    var a string
    fmt.Scan(&s1, &a)
    
    
    count, _ := strconv.Atoi(a)
    for i := 0; i < count; i++{
        *p2 += s1
    }
    fmt.Println(s2)
}