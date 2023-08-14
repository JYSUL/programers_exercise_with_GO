func solution(my_string string, s int, e int) string {
    var answer string = ""
    answer += my_string[:s]
    for i := e; i > s-1; i-- {
        answer += string(my_string[i])
    }
    answer += my_string[e+1:]
    return answer
}