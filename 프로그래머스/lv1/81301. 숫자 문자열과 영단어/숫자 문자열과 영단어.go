import (
    //"fmt"
    "strings"
    "strconv"
)
func solution(s string) int {
    alpha_num := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    for index, alpha := range alpha_num {
        s = strings.Replace(s, alpha,strconv.Itoa(index),-1)
    }
    
    //fmt.Println(s)
    
    answer, err := strconv.Atoi(s)
    if err != nil {
        panic("strconv error, answer, Atoi")
    }
    
    return answer
}