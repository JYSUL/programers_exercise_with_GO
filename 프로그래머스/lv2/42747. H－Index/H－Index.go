import (
    "sort"
    //"fmt"
)
func solution(citations []int) int {
    var answer int
    sort.Sort(sort.Reverse(sort.IntSlice(citations)))
    for index, citation := range citations {
        //fmt.Println(index+1, citation)
        if index + 1 <= citation {
            answer = index + 1
        } else {
            break
        }
    }
    return answer
}