func solution(n int) string {
    var 수박 []string = []string{"수", "박"}
    var answer string = ""
    for i := 0; i < n; i++ {
        answer += 수박[i%2]
    }
    return answer
}