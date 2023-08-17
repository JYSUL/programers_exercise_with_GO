import (
    //"fmt"
    "strings"
    "sort"
)
func solution(myString string) []string {
    var answer []string
    answer = strings.Fields(strings.Replace(myString,"x"," ",-1))
    sort.Strings(answer)
    return answer
}