func solution(num int, n int) int {
    var answer_flag int = 0
    if num % n == 0 {
        answer_flag = 1
    }
    return answer_flag
}