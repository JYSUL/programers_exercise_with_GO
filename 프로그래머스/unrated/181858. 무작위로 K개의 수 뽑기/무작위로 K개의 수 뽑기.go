//import "fmt"
func solution(arr []int, k int) []int {
    flag_table := make(map[int]bool)
    var answer []int
    for _, num := range arr {
        if !flag_table[num] {
            flag_table[num] = true
            answer = append(answer, num)
            if len(answer) >= k {
                break
            }
        }
    }
    for i := len(answer); i < k; i ++ {
        answer = append(answer, -1)
    }
    return answer
}