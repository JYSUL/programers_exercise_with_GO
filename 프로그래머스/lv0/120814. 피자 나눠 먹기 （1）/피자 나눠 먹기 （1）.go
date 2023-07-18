func solution(n int) int {
    var answer int = n / 7
    if n % 7 != 0 {
        answer++
    }
    return answer
}