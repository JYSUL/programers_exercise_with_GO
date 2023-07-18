func solution(slice int, n int) int {
    var answer int
    if n % slice == 0 {
        answer = n / slice
    } else {
        answer = n / slice + 1
    }
    return answer
}