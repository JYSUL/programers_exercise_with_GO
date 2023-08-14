func solution(num_list []int, n int) []int {
    var answer []int
    answer = append(answer, num_list[n:]...)
    answer = append(answer, num_list[:n]...)
    return answer
}