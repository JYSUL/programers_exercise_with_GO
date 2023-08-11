func solution(a int, d int, included []bool) int {
    var answer int = 0
    for index, flag := range included {
        if flag {
            answer += a + d * index
        }        
    }
    return answer
}