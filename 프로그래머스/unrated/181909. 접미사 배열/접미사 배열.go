import (
    "sort"
)
func solution(my_string string) []string {
    var answer_slice []string
    for i := 0; i < len(my_string); i++ {
        answer_slice = append(answer_slice, my_string[i:])
    }
    sort.Strings(answer_slice)
    return answer_slice
}