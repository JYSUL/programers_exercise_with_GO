//import "fmt"
func solution(s string) []int {
    var answer []int
    index_table := make(map[int32]int)
    for index, str := range s {
        last_index, flag := index_table[str]
        if flag {
            answer = append(answer, index - last_index)
        } else {
            answer = append(answer, -1)
        }
        index_table[str] = index
    }
    return answer
}