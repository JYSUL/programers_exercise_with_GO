package main

import "fmt"

func main() {
    var s1 string
    var s2 string
    var p2 *string = &s2
    
    fmt.Scan(&s1)
    
    for _, str := range(s1) {
        if str >= 97 {            
            *p2 += string(str-32)
        } else {
            *p2 += string(str+32)
        }
    }
    fmt.Println(s2)
}