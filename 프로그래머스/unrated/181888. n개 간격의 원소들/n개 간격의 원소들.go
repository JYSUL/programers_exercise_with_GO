func solution(num_list []int, n int) []int {
    var answer_slice []int
    for i := 0; i < len(num_list); i += n {
        answer_slice = append(answer_slice, num_list[i])
    }
    return answer_slice
}