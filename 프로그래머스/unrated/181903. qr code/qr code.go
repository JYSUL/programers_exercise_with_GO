func solution(q int, r int, code string) string {
    var answer string
    for i := r; i<len(code); i += q {
        answer += string(code[i])
    }
    return answer
}