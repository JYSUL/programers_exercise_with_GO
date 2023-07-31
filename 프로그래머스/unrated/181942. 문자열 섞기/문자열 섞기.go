func solution(str1 string, str2 string) string {
    var answer_str string = ""
    var p *string = &answer_str
    for i := 0; i < len(str1); i++ {
        *p += string(str1[i]) + string(str2[i])
    }
    return answer_str
}