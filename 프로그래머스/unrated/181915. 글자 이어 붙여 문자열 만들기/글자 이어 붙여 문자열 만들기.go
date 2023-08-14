func solution(my_string string, index_list []int) string {
    var answer string = ""
    for _, index := range index_list {
        answer += string(my_string[index])
    }
    return answer
}