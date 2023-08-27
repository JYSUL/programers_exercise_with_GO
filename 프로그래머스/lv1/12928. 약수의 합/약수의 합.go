func solution(n int) int {
    var answer int = 0
    for i := 1; i <= n; i++ {
        if n % i == 0 {
            answer += i
        }
    }
    return answer
}