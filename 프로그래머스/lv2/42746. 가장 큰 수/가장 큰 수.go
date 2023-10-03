//import "fmt"

import (
    "sort"
    "strconv"
    "strings"
)

func solution(numbers []int) string {
    var str_numbers []string
    for _, number := range numbers {
        str_numbers = append(str_numbers, strconv.Itoa(number))
    }
    //fmt.Println(str_numbers)
    sort.Slice(str_numbers, func(i, j int) bool {
        return str_numbers[i]+str_numbers[j] > str_numbers[j] + str_numbers[i]       
    })
    
    if str_numbers[0] == "0" {
        return "0"
    }
    
    //fmt.Println(str_numbers)
    return strings.Join(str_numbers, "")
}