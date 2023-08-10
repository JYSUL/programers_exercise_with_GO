func solution(number int, n int, m int) int {
    var answer int = 0
    if number % n == 0 && number % m == 0 {
        answer = 1
    }
    return answer
}