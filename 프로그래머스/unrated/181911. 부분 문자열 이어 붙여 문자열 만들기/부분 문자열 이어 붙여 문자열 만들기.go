func solution(my_strings []string, parts [][]int) string {
    var answer string = ""
    for i := 0 ; i < len(parts); i++ {
        answer += my_strings[i][parts[i][0]:parts[i][1]+1]
    }
    return answer
}