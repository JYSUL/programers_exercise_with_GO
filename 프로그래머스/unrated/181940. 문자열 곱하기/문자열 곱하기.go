func solution(my_string string, k int) string {
    var answer_str string
    var p_answer_str *string = &answer_str
    for i := 0; i<k; i++ {
        *p_answer_str += my_string
    }
    return answer_str
}