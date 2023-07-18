func solution(n int) int {
    var answer int = 0
    
    for i := 0; i < n+1; i = i + 2 {
        answer = answer + i
    }
    return answer
}