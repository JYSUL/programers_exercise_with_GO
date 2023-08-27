func solution(x int, n int) []int64 {
    var answer []int64
    for i := 1; i <= n; i++ {
        answer = append(answer, int64(i*x))
    }
    return answer
}