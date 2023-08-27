func solution(n int) int {
    answer := 2
    for n % answer != 1 {
        answer++
    }
    return answer
}