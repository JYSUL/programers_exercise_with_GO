//import "fmt"

import (
    "strconv"
    "strings"
)

//reverse_string is return reversed string
func reverse_string(str string) string {
    runes := []rune(str)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}


func solution(numbers []int64) []int64 {
    var result []int64
    for _, number := range numbers {
        //짝수면 +1, 홀수면 2진 숫자의 뒤에서부터 확인하면서 최초의 01을 10으로 바꾼다
        if number % 2 == 0 {
            result = append(result, number + 1)
        } else {
            bin_num := "0" + strconv.FormatInt(number, 2)
            bin_num = reverse_string(strings.Replace(reverse_string(bin_num), "10", "01", 1))
            new_num, err := strconv.ParseInt(bin_num, 2, 64)
            if err != nil {
                panic("strconv.ParseInt")
            }
            result = append(result, new_num)
        }
    }
    return result
}