func solution(names []string) []string {
    var answer []string
    for i := 0; i < len(names); i += 5 {
        answer = append(answer, names[i])
    }
    return answer
}