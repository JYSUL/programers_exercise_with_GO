func solution(s1 []string, s2 []string) int {
    var answer int = 0
    for _, s1_op := range s1 {
        for _, s2_op := range s2 {
            if s1_op == s2_op {
                answer++
                break
            }
        }
    }
    return answer
}