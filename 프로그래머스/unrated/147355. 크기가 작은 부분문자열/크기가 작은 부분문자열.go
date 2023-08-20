func solution(t string, p string) int {
    var answer int
    for i := 0; i < len(t) - len(p) + 1; i++ {
        if t[i:i+len(p)] <= p {
            answer++
        }
    }
    return answer
}