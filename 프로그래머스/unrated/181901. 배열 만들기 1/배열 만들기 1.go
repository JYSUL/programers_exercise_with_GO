func solution(n int, k int) []int {
    var answer []int
    for i := k; i <= n;i += k {
        answer = append(answer, i)
    } 
    return answer
}