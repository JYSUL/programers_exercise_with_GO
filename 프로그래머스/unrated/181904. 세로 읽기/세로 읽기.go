func solution(my_string string, m int, c int) string {
    var answer string = ""
    for i := c-1; i< len(my_string); i += m {
        answer += string(my_string[i])
    }
    return answer
}