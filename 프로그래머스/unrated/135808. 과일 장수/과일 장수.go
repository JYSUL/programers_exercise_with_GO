import (
    "sort"
    //"fmt"
)
func solution(k int, m int, score []int) int {
    var answer int
    sort.Sort(sort.Reverse(sort.IntSlice(score)))
    for i := m-1; i < len(score); i += m {
        answer += score[i] * m
        //fmt.Println(score)
        //fmt.Println(score[i] * m)
    }
    return answer
}